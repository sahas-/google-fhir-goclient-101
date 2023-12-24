package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
	r4pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
)

func main() {
	files := readFolder(INPUT_FOLDER)
	reorderedFiles := reorderFiles(files, resourceOrder)
	for _, file := range reorderedFiles {
		filepath := filepath.Join(INPUT_FOLDER, file.Name())
		fmt.Println("Processing ", file.Name())
		parseNdjsonData(filepath)
	}

}
func readFolder(inputFolderPath string) []fs.FileInfo {
	files, err := ioutil.ReadDir(inputFolderPath)
	if err != nil {
		panic(err)
	}
	return (files)
}
func parseNdjsonData(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	// Iterate over the lines in the file
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		um, err := jsonformat.NewUnmarshaller(timeZone, fhirversion.R4)
		if err != nil {
			panic(err)
		}
		unmarshalled, err := um.Unmarshal(line)
		if err != nil {
			panic(err)
		}

		contained := unmarshalled.(*r4pb.ContainedResource)
		oneOfResource := contained.GetOneofResource()
		switch r := oneOfResource.(type) {
		case *r4pb.ContainedResource_Organization:
			org := contained.GetOrganization().Id.Value
			postToFHIRServer(line, org, ORG_RESOURCE)
		case *r4pb.ContainedResource_Location:
			loc := contained.GetLocation().Id.Value
			postToFHIRServer(line, loc, LOCATION_RESOURCE)
		case *r4pb.ContainedResource_Practitioner:
			practitioner := contained.GetPractitioner().Id.Value
			postToFHIRServer(line, practitioner, PRACTITIONER_RESOURCE)
		case *r4pb.ContainedResource_PractitionerRole:
			practitionerRole := contained.GetPractitionerRole().Id.Value
			postToFHIRServer(line, practitionerRole, PRACTITIONER_ROLE_RESOURCE)
		case *r4pb.ContainedResource_Condition:
			condition := contained.GetCondition().Id.Value
			postToFHIRServer(line, condition, CONDITION_RESOURCE)
		case *r4pb.ContainedResource_Patient:
			patient := contained.GetPatient().Id.Value
			postToFHIRServer(line, patient, PATIENT_RESOURCE)
		case *r4pb.ContainedResource_AllergyIntolerance:
			allergy := contained.GetAllergyIntolerance().Id.Value
			postToFHIRServer(line, allergy, ALLERGY_RESOURCE)
		case *r4pb.ContainedResource_Medication:
			medication := contained.GetMedication().Id.Value
			postToFHIRServer(line, medication, MEDICATION_RESOURCE)
		case *r4pb.ContainedResource_Device:
			device := contained.GetDevice().Id.Value
			postToFHIRServer(line, device, DEVICE_RESOURCE)
		case *r4pb.ContainedResource_CarePlan:
			carePlan := contained.GetCarePlan().Id.Value
			postToFHIRServer(line, carePlan, CAREPLAN_RESOURCE)
		case *r4pb.ContainedResource_CareTeam:
			careTeam := contained.GetCareTeam().Id.Value
			postToFHIRServer(line, careTeam, CARETEAM_RESOURCE)
		case *r4pb.ContainedResource_MedicationAdministration:
			medicationAdministration := contained.GetMedicationAdministration().Id.Value
			postToFHIRServer(line, medicationAdministration, MEDICATION_ADMINISTIONATION_RESOURCE)
		case *r4pb.ContainedResource_MedicationStatement:
			medicationStatement := contained.GetMedicationStatement().Id.Value
			postToFHIRServer(line, medicationStatement, MEDICATION_ADMINISTIONATION_RESOURCE)
		case *r4pb.ContainedResource_MedicationRequest:
			medicationReq := contained.GetMedicationRequest().Id.Value
			postToFHIRServer(line, medicationReq, MEDICATION_REQ_RESOURCE)
		case *r4pb.ContainedResource_Observation:
			observation := contained.GetObservation().Id.Value
			postToFHIRServer(line, observation, OBSERVATION_RESOURCE)
		case *r4pb.ContainedResource_DiagnosticReport:
			diagnosticReport := contained.GetDiagnosticReport().Id.Value
			postToFHIRServer(line, diagnosticReport, DIAGNOSTIC_REPORT_RESOURCE)
		case *r4pb.ContainedResource_Immunization:
			immunization := contained.GetImmunization().Id.Value
			postToFHIRServer(line, immunization, IMMUNIZATION_RESOURCE)
		case *r4pb.ContainedResource_Encounter:
			encounter := contained.GetEncounter().Id.Value
			postToFHIRServer(line, encounter, ENCOUNTER_RESOURCE)
		case *r4pb.ContainedResource_Claim:
			claim := contained.GetClaim().Id.Value
			postToFHIRServer(line, claim, CLAIM_RESOURCE)
		case *r4pb.ContainedResource_ExplanationOfBenefit:
			explanationOfBenefit := contained.GetExplanationOfBenefit().Id.Value
			postToFHIRServer(line, explanationOfBenefit, EXPLANATION_OF_BENEFIT_RESOURCE)
		case *r4pb.ContainedResource_DocumentReference:
			docRef := contained.GetDocumentReference().Id.Value
			postToFHIRServer(line, docRef, DOCUMENT_REFERENCE_RESOURCE)
		case *r4pb.ContainedResource_Provenance:
			provenance := contained.GetProvenance().Id.Value
			postToFHIRServer(line, provenance, PROVENANCE_RESOURCE)
		case *r4pb.ContainedResource_SupplyDelivery:
			supplyChain := contained.GetSupplyDelivery().Id.Value
			postToFHIRServer(line, supplyChain, SUPPLYDELIVERY_RESOURCE)
		case *r4pb.ContainedResource_ImagingStudy:
			imagingStudy := contained.GetImagingStudy().Id.Value
			postToFHIRServer(line, imagingStudy, IMAGINGSTUDY_RESOURCE)
		case *r4pb.ContainedResource_Procedure:
			procedure := contained.GetProcedure().Id.Value
			postToFHIRServer(line, procedure, PROCEDURE_RESOURCE)
		default:
			fmt.Printf("Other resource type: %T\n", r)
		}
	}
}
func postToFHIRServer(body []byte, id string, resource string) {
	url := FHIR_SERVER + resource + "/" + id
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Add("x-api-key", API_KEY)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	// fmt.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	// Check for successful update
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		fmt.Println("Created ", resource)
	} else {
		fmt.Println("Error updating:", resp.Status)
		panic(resp.StatusCode)
	}

}
