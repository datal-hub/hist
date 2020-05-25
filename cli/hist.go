package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"

	"hist/models"

	"github.com/urfave/cli"
)

var maxProcs int
var countAggregators int

//description run the application from the command line for compute hist for symbols in files of directory
var Hist = cli.Command{
	Name:        "hist",
	Usage:       "Count hist ascii symbols from file",
	Description: ``,
	Action:      hist,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:        "max_procs, p",
			Value:       8,
			Usage:       "Define value for GOMAXPROCS",
			Destination: &maxProcs},
		cli.IntFlag{
			Name:        "aggregators, a",
			Value:       100,
			Usage:       "Count of routines for aggregating hist from files",
			Destination: &countAggregators},
		cli.StringFlag{
			Name:        "path_to_dir, path",
			Value:       "./examples",
			Usage:       "Path to directory with files",
			Destination: &dirPath},
	},
}

var reportCountMutex sync.Mutex
var preComputeHistMutex sync.Mutex
var aggregatorsGroup sync.WaitGroup
var fileReadersGroup sync.WaitGroup

// fileToHist is reading text file and counting ascii symbols to AsciiHist
func fileToHist(filePath string) models.AsciiHist {
	result := models.AsciiHist{}
	defer fileReadersGroup.Done()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Print(err.Error())
		return result
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err.Error())
		return result
	}
	result.FromByteSlice(content)
	return result
}

// This function aggregating hist from routines that running for all of files
// and putting results to preComputeHist with index of current aggregator routine
func aggregator(number int, preComputeHist []models.AsciiHist, histChan chan models.AsciiHist) {
	tmpResHist := models.AsciiHist{}
	defer aggregatorsGroup.Done()
	for {
		select {
		case curHist, ok := <-histChan:
			if ok {
				tmpResHist.Add(curHist)
			} else {
				preComputeHistMutex.Lock()
				preComputeHist[number] = tmpResHist
				preComputeHistMutex.Unlock()
				return
			}
		}
	}
}

// This function launching routine for each file in directory that count hist.
// Results each routine sending on aggregators that putting several common results to preComputeHist
// and after this counting final result.
func histFromDir(dirPath string, numprocs int, numAggregators int) (models.AsciiHist, error) {
	runtime.GOMAXPROCS(numprocs)
	histChan := make(chan models.AsciiHist, numAggregators)
	reportCount := 0
	resHist := models.AsciiHist{}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
		return resHist, err
	}
	if numAggregators > len(files) {
		numAggregators = len(files)
	}
	var preComputeHist = make([]models.AsciiHist, numAggregators)
	for _, file := range files {
		fileReadersGroup.Add(1)
		filePath := fmt.Sprintf("%s/%s", dirPath, file.Name())
		go func(pathToFile string) {
			hist := fileToHist(pathToFile)
			reportCountMutex.Lock()
			histChan <- hist
			reportCount++
			if reportCount == len(files) {
				close(histChan)
			}
			reportCountMutex.Unlock()
		}(filePath)
	}
	for i := 0; i < numAggregators; i++ {
		aggregatorsGroup.Add(1)
		go aggregator(i, preComputeHist, histChan)
	}
	aggregatorsGroup.Wait()
	fileReadersGroup.Wait()
	for i := 0; i < numAggregators; i++ {
		resHist.Add(preComputeHist[i])
	}
	return resHist, nil
}

func hist(c *cli.Context) error {
	hist, err := histFromDir(dirPath, maxProcs, countAggregators)
	if err != nil {
		fmt.Printf("error msg: %s", err.Error())
		return err
	}
	fmt.Print(hist)
	return nil
}
