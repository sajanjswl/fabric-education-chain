package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/jung-kurt/gofpdf"
)

const ipfsURL = "ipfs:5001"

const (
	Subject1Name = "Physics"
	Subject1Code = "S0001"
	Subject2Name = "Chemistry"
	Subject2Code = "S0002"
	Subject3Name = "Mathmaitcs"
	Subject3Code = "S0003"
	Faculty1Code = "F0001"
	Faculty2Code = "F0002"
	Faculty3Code = "F0003"
)

// A stringArray is an array of strings that has been unmarshalled from a JSON
// property that could be either a string or an array of string
type stringArray []string

// SmartContract provides functions for managing a student
type SmartContract struct {
	contractapi.Contract
}

// basic student
type Student struct {
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	Branch             string    `json:"branch"`
	RegistrationNumber string    `json:"registrationNumber"`
	BloodGroup         string    `json:"bloodGroup"`
	MobileNumber       string    `json:"mobileNumber"`
	Address            string    `json:"address"`
	Subjects           []Subject `json:"subjects"`
	ReportHash         string    `json:"reportHash"`
}

type Subject struct {
	SubjectCode  string  `json:"subjectCode"`
	SubjectName  string  `json:"subjectName"`
	SubjectMarks float64 `json:"subjectMarks"`
	FacultyCode  string  `json:"facultyCode"`
}

type Professor struct {
	FacultyCode string      `json:"facultyCode"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Subjects    stringArray `json:"subject"`
}

// func (sa *stringArray) UnmarshalJSON(data []byte) error {
// 	if len(data) > 0 {
// 		switch data[0] {
// 		case '"':
// 			var s string
// 			if err := json.Unmarshal(data, &s); err != nil {
// 				return err
// 			}
// 			*sa = []string{s}
// 		case '[':
// 			var s []string
// 			if err := json.Unmarshal(data, &s); err != nil {
// 				return err
// 			}
// 			*sa = s
// 		}
// 	}
// 	return nil
// }

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key     string `json:"key"`
	Student *Student
}

// InitLedger adds a base set of student to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	professors := []Professor{
		{
			FacultyCode: Faculty1Code,
			Subjects:    stringArray{Subject1Code},
			FirstName:   "Amit",
			LastName:    "Jaiswal",
		},
		{
			FacultyCode: Faculty2Code,
			Subjects:    stringArray{Subject2Code},
			FirstName:   "Sujit",
			LastName:    "Jaiswal",
		},
		{
			FacultyCode: Faculty3Code,
			Subjects:    stringArray{Subject3Code},
			FirstName:   "Kapil",
			LastName:    "Jaiswal",
		},
	}

	for _, professor := range professors {
		professorAsBytes, _ := json.Marshal(professor)
		if err := ctx.GetStub().PutState(professor.FacultyCode, professorAsBytes); err != nil {
			return fmt.Errorf("Failed to put to world state. %s", professor.FacultyCode)
		}
	}

	subjects := []Subject{{FacultyCode: Faculty1Code, SubjectCode: Subject1Code, SubjectName: Subject1Name, SubjectMarks: 80},
		{FacultyCode: Faculty2Code, SubjectCode: Subject2Code, SubjectName: Subject2Name, SubjectMarks: 90},
		{FacultyCode: Faculty3Code, SubjectCode: Subject3Code, SubjectName: Subject3Name, SubjectMarks: 100},
	}

	students := []Student{
		{FirstName: "Sajan", LastName: "Jaiswal", Branch: "CSE", RegistrationNumber: "1816128", BloodGroup: "A+",
			MobileNumber: "+917064274923", Address: "White House, Motihari, Bihar", Subjects: subjects,
		},
		{FirstName: "Abhishek", LastName: "Jaiswal", Branch: "CSE", RegistrationNumber: "1816129",
			BloodGroup: "B+", MobileNumber: "+918210791275", Address: "MidLand,Dimapur", Subjects: subjects,
		},
	}

	for _, student := range students {
		studentAsBytes, _ := json.Marshal(student)
		if err := ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes); err != nil {
			return fmt.Errorf("Failed to put to world state. %s", student.RegistrationNumber)
		}
	}

	return nil
}

// RegisterStudent adds a new Student to the world state with given details
func (s *SmartContract) RegisterStudent(ctx contractapi.TransactionContextInterface, registrationNumber, firstName, lastName, branch, bloodGroup, mobileNumber, address string) error {

	// subjects := stringArray{"Physics", "chemistry", "Math"}
	subjects := []Subject{
		{FacultyCode: Faculty1Code,
			SubjectCode: Subject1Code,
			SubjectName: Subject1Name,
		},
		{FacultyCode: Faculty2Code,
			SubjectCode: Subject2Code,
			SubjectName: Subject2Name,
		},
		{FacultyCode: Faculty3Code,
			SubjectCode: Subject3Code,
			SubjectName: Subject3Name,
		},
	}
	student := Student{
		FirstName:          firstName,
		LastName:           lastName,
		Branch:             branch,
		RegistrationNumber: registrationNumber,
		BloodGroup:         bloodGroup,
		MobileNumber:       mobileNumber,
		Address:            address,
		Subjects:           subjects,
	}

	studentAsBytes, _ := json.Marshal(&student)
	return ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes)
}

// QueryStudent returns the student stored in the world state with given id
func (s *SmartContract) QueryStudent(ctx contractapi.TransactionContextInterface, registrationNumber string) (*Student, error) {
	cidAsBytes, err := ctx.GetStub().GetState(registrationNumber)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", registrationNumber)
	}
	if cidAsBytes == nil {
		return nil, fmt.Errorf("student %s does not exist", registrationNumber)
	}
	student := new(Student)

	err = json.Unmarshal(cidAsBytes, student)

	if err != nil {
		return nil, fmt.Errorf("failed to parse json %s", err)
	}
	return student, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
func (s *SmartContract) UpdateMobileNumber(ctx contractapi.TransactionContextInterface, registrationNumber string, mobileNumber string) error {
	student, err := s.QueryStudent(ctx, registrationNumber)

	if err != nil {
		return err
	}
	student.MobileNumber = mobileNumber

	// Converting the map into JSON object
	studentAsBytes, _ := json.Marshal(student)

	return ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes)

}

func (s *SmartContract) RewardMarks(ctx contractapi.TransactionContextInterface, facultyCode, studentRegistrationNumber, subjectCode string, marks float64) error {

	studentAsBytes, err := ctx.GetStub().GetState(studentRegistrationNumber)
	if err != nil {
		return fmt.Errorf("Failed to read from world state. %s", studentRegistrationNumber)
	}
	if studentAsBytes == nil {
		return fmt.Errorf("student %s does not exist", studentRegistrationNumber)
	}
	student := new(Student)
	if err = json.Unmarshal(studentAsBytes, student); err != nil {
		return fmt.Errorf("failed to parse json %s", err)
	}

	for i, subject := range student.Subjects {
		if subject.FacultyCode == facultyCode && subject.SubjectCode == subjectCode {
			if subject.SubjectMarks == 0 {
				student.Subjects[i].SubjectMarks = marks
			} else {
				return fmt.Errorf("Marks Already rewarded")
			}

		}

	}

	// Converting the map into JSON object
	studentAsBytes, _ = json.Marshal(student)

	return ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes)
}

func (s *SmartContract) GenrateReport(ctx contractapi.TransactionContextInterface, registrationNumber string) (string, error) {
	studentAsBytes, err := ctx.GetStub().GetState(registrationNumber)
	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", registrationNumber)
	}
	if studentAsBytes == nil {
		return "", fmt.Errorf("student %s does not exist", registrationNumber)
	}
	student := new(Student)

	if err = json.Unmarshal(studentAsBytes, student); err != nil {
		return "", fmt.Errorf("failed to parse json %s", err)
	}

	if student.ReportHash != "" {

		return student.ReportHash, nil
	}

	pdf := newReport(student.FirstName+" "+student.LastName, student.RegistrationNumber)
	if pdf.Err() {
		return "", fmt.Errorf("Failed creating PDF report: %s\n", pdf.Error())
	}

	subectNameArray := []string{Subject1Name, Subject2Name, Subject3Name, "Percentage"}

	if pdf = header(pdf, subectNameArray); pdf.Err() {
		return "", fmt.Errorf("Failed creating PDF report: %s\n", pdf.Error())
	}

	aggregate := (student.Subjects[0].SubjectMarks + student.Subjects[1].SubjectMarks + student.Subjects[2].SubjectMarks) / 3
	subjectMarksArray := [][]string{{
		fmt.Sprintf("%.2f", student.Subjects[0].SubjectMarks),
		fmt.Sprintf("%.2f", student.Subjects[1].SubjectMarks),
		fmt.Sprintf("%.2f", student.Subjects[2].SubjectMarks),
		fmt.Sprintf("%.2f", aggregate),
	},
	}

	if pdf = table(pdf, subjectMarksArray); pdf.Err() {
		return "", fmt.Errorf("Failed creating PDF report: %s\n", pdf.Error())
	}

	pipeReader, pipeWriter := io.Pipe()

	go func() {
		if err := pdf.OutputAndClose(pipeWriter); err != nil {
			log.Printf("Faile to wirte File to io writer %s", err)
		}
		pipeWriter.Close()

	}()

	reportHash, err := s.uploadReportOnIPFS(pipeReader)
	if err != nil {
		return "", err
	}
	student.ReportHash = reportHash
	studentAsBytes, _ = json.Marshal(student)
	if err := ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes); err != nil {
		return "", err
	}
	return reportHash, nil
}

func (c *SmartContract) uploadReportOnIPFS(pipeReader *io.PipeReader) (string, error) {

	sh := shell.NewShell(ipfsURL)

	reportHash, err := sh.Add(pipeReader)
	if err != nil {
		return "", fmt.Errorf("Faile to upload file to ipfs %s", err)
	}

	return reportHash, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create student chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting student chaincode: %s", err.Error())
	}
}

func table(pdf *gofpdf.Fpdf, tbl [][]string) *gofpdf.Fpdf {
	// Reset font and fill color.
	pdf.SetFont("Times", "", 16)
	pdf.SetFillColor(255, 255, 255)

	// Every column gets aligned according to its contents.
	align := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range tbl {
		for i, str := range line {
			// Again, we need the CellFormat() method to create a visible border around the cell. We also use the alignStr parameter here to print the cell content either left-aligned or right-aligned.
			pdf.CellFormat(40, 7, str, "1", 0, align[i], false, 0, "")
		}
		pdf.Ln(-1)
	}
	return pdf
}

func newReport(name, registrationNumber string) *gofpdf.Fpdf {

	pdf := gofpdf.New("L", "mm", "Letter", "")

	pdf.AddPage()

	pdf.SetFont("Times", "B", 28)
	pdf.Cell(80, 10, "Grade Sheet")
	pdf.Ln(12)

	pdf.SetFont("Times", "", 20)
	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
	pdf.Ln(14)

	pdf.SetFont("Times", "", 18)
	pdf.Cell(80, 10, name)
	pdf.Ln(16)
	pdf.SetFont("Times", "", 18)
	pdf.Cell(80, 10, registrationNumber)
	pdf.Ln(18)

	return pdf
}

func header(pdf *gofpdf.Fpdf, hdr []string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	for _, str := range hdr {

		pdf.CellFormat(40, 7, str, "1", 0, "", true, 0, "")
	}

	pdf.Ln(-1)
	return pdf
}
