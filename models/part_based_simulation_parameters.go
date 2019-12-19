// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartBasedSimulationParameters part based simulation parameters
// swagger:model PartBasedSimulationParameters
type PartBasedSimulationParameters struct {

	// should be a number between 0.5 and 1.5
	BladeCrashThreshold float64 `json:"bladeCrashThreshold,omitempty"`

	// Id of build file being simulated, mutually exclusive with simulationParts
	BuildFileID int32 `json:"buildFileId,omitempty"`

	// cutoff method
	// Required: true
	CutoffMethod *string `json:"cutoffMethod"`

	// detect blade crash
	// Required: true
	DetectBladeCrash *bool `json:"detectBladeCrash"`

	// detect support failure
	DetectSupportFailure bool `json:"detectSupportFailure,omitempty"`

	// a value that is used to scale the after cutoff simulated distortion values
	// Maximum: 5
	// Minimum: -5
	DistortionAfterCutoffScaleFactor *float64 `json:"distortionAfterCutoffScaleFactor,omitempty"`

	// Array of distortion after cutoff scale factor values, each one will produce a separate distortion after cutoff output
	DistortionAfterCutoffScaleFactorValues []*float64 `json:"distortionAfterCutoffScaleFactorValues"`

	// a value that is used to scale the simulated distortion value
	// Maximum: 5
	// Minimum: -5
	DistortionScaleFactor *float64 `json:"distortionScaleFactor,omitempty"`

	// Array of distortion scale factor values, each one will produce a separate distortion output
	DistortionScaleFactorValues []*float64 `json:"distortionScaleFactorValues"`

	// elastic modulus
	// Required: true
	ElasticModulus *float64 `json:"elasticModulus"`

	// if true, contour scan vectors is used in scan pattern simulation
	EnableContours bool `json:"enableContours,omitempty"`

	// generate support voxels
	// Required: true
	GenerateSupportVoxels *bool `json:"generateSupportVoxels"`

	// hardening factor
	// Required: true
	// Maximum: 0.5
	// Minimum: 0
	HardeningFactor *float64 `json:"hardeningFactor"`

	// if true, the on-plate stress output file will include the on-plate strain
	IncludeOnPlateStrainOutput bool `json:"includeOnPlateStrainOutput,omitempty"`

	// load stepping type
	LoadSteppingType string `json:"loadSteppingType,omitempty"`

	// load steps
	// Maximum: 200
	// Minimum: 1
	LoadSteps int32 `json:"loadSteps,omitempty"`

	// Must be between 0.00015 to 0.002 meters, Must be greater than minimumWallThickness
	// Maximum: 0.01
	// Minimum: 2e-05
	MaximumThickWallThickness float64 `json:"maximumThickWallThickness,omitempty"`

	// The maximum distance between thinwalls
	// Maximum: 0.01
	// Minimum: 2e-05
	MaximumThinWallDistance float64 `json:"maximumThinWallDistance,omitempty"`

	// Distance to move the part off the base plate for supports, Must be between 0 to 0.005 meters
	// Required: true
	// Maximum: 0.005
	// Minimum: 0
	MinimumSupportHeight *float64 `json:"minimumSupportHeight"`

	// Must be between 0.00005 to 0.0003 meters, Must be less than maximumWallThickness
	// Maximum: 0.01
	// Minimum: 2e-05
	MinimumThickWallThickness float64 `json:"minimumThickWallThickness,omitempty"`

	// output displacement after cutoff
	// Required: true
	OutputDisplacementAfterCutoff *bool `json:"outputDisplacementAfterCutoff"`

	// if true, mechanics solver output will include a zip file with the stress / distortion state at the end of each voxel layer
	// Required: true
	OutputLayerVtk *bool `json:"outputLayerVtk"`

	// if true, mechanics solver output will include files related to passing mesh and inital state to ANSYS Workbench
	// Required: true
	OutputMapdl *bool `json:"outputMapdl"`

	// if true, a VTK file of the support structure will be created
	OutputSupportsVtk bool `json:"outputSupportsVtk,omitempty"`

	// should be a number between 0.01 and 1.0
	PartFailureThreshold float64 `json:"partFailureThreshold,omitempty"`

	// part file location
	PartFileLocation string `json:"partFileLocation,omitempty"`

	// if true, a predistorted STL file will be created using the distortion simulated by the mechanics solver
	PerformDistortionCompensation bool `json:"performDistortionCompensation,omitempty"`

	// if true, a predistorted STL file will be created using the distortion after cutoff simulated by the mechanics solver
	PerformDistortionCompensationAfterCutoff bool `json:"performDistortionCompensationAfterCutoff,omitempty"`

	// perform support only cutoff
	// Required: true
	PerformSupportOnlyCutoff *bool `json:"performSupportOnlyCutoff"`

	// perform support optimization
	// Required: true
	PerformSupportOptimization *bool `json:"performSupportOptimization"`

	// poisson ratio
	// Required: true
	PoissonRatio *float64 `json:"poissonRatio"`

	// The offset between the positioned part and the origin in x direction
	PositionedPartOffsetX float64 `json:"positionedPartOffsetX,omitempty"`

	// The offset between the positioned part and the origin in y direction
	PositionedPartOffsetY float64 `json:"positionedPartOffsetY,omitempty"`

	// The offset between the positioned part and the origin in z direction
	PositionedPartOffsetZ float64 `json:"positionedPartOffsetZ,omitempty"`

	// List of parts to simulate (current limit is one part, imposed by server)
	SimulationParts []*SimulationPart `json:"simulationParts"`

	// strain scaling factor
	// Required: true
	StrainScalingFactor *float64 `json:"strainScalingFactor"`

	// should be a number between 0.01 and 1.0
	StrainWarningThreshold float64 `json:"strainWarningThreshold,omitempty"`

	// stress mode
	// Required: true
	StressMode *string `json:"stressMode"`

	// Must be between 1 to 89 degrees
	// Required: true
	// Maximum: 89
	// Minimum: 1
	SupportAngle *float64 `json:"supportAngle"`

	// Multiplier for support calculations, Must be between 0.1 to 10
	// Required: true
	// Maximum: 10
	// Minimum: 0.1
	SupportFactorOfSafety *float64 `json:"supportFactorOfSafety"`

	// should be a number between 0.01 and 1.0
	SupportFailureThreshold float64 `json:"supportFailureThreshold,omitempty"`

	// support file location
	SupportFileLocation string `json:"supportFileLocation,omitempty"`

	// type of support used for simulation.
	// Required: true
	SupportType *string `json:"supportType"`

	// support yield strength
	// Required: true
	SupportYieldStrength *float64 `json:"supportYieldStrength"`

	// support yield strength ratio
	// Required: true
	SupportYieldStrengthRatio *float64 `json:"supportYieldStrengthRatio"`

	// The distance between thickwall supports
	// Maximum: 0.01
	// Minimum: 2e-05
	ThickWallDistance float64 `json:"thickWallDistance,omitempty"`

	// The width of the thinwall supports
	// Maximum: 0.0005
	// Minimum: 2e-05
	ThinWallThickness float64 `json:"thinWallThickness,omitempty"`

	// the amount of sub voxels per voxel in each direction (x,y,z)
	// Maximum: 10
	// Minimum: 1
	VoxelSampleRate int32 `json:"voxelSampleRate,omitempty"`

	// Must be between 0.00002 to 0.002 meters
	// Required: true
	// Maximum: 0.002
	// Minimum: 2e-05
	VoxelSize *float64 `json:"voxelSize"`
}

// Validate validates this part based simulation parameters
func (m *PartBasedSimulationParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCutoffMethod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDetectBladeCrash(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDistortionAfterCutoffScaleFactor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDistortionAfterCutoffScaleFactorValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDistortionScaleFactor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDistortionScaleFactorValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateElasticModulus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGenerateSupportVoxels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHardeningFactor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLoadSteppingType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLoadSteps(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaximumThickWallThickness(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaximumThinWallDistance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinimumSupportHeight(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinimumThickWallThickness(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputDisplacementAfterCutoff(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputLayerVtk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputMapdl(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePerformSupportOnlyCutoff(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePerformSupportOptimization(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePoissonRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSimulationParts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStrainScalingFactor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStressMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportFactorOfSafety(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportYieldStrength(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportYieldStrengthRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThickWallDistance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThinWallThickness(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVoxelSampleRate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVoxelSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var partBasedSimulationParametersTypeCutoffMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Legacy","Instantaneous"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partBasedSimulationParametersTypeCutoffMethodPropEnum = append(partBasedSimulationParametersTypeCutoffMethodPropEnum, v)
	}
}

const (
	// PartBasedSimulationParametersCutoffMethodLegacy captures enum value "Legacy"
	PartBasedSimulationParametersCutoffMethodLegacy string = "Legacy"
	// PartBasedSimulationParametersCutoffMethodInstantaneous captures enum value "Instantaneous"
	PartBasedSimulationParametersCutoffMethodInstantaneous string = "Instantaneous"
)

// prop value enum
func (m *PartBasedSimulationParameters) validateCutoffMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partBasedSimulationParametersTypeCutoffMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartBasedSimulationParameters) validateCutoffMethod(formats strfmt.Registry) error {

	if err := validate.Required("cutoffMethod", "body", m.CutoffMethod); err != nil {
		return err
	}

	// value enum
	if err := m.validateCutoffMethodEnum("cutoffMethod", "body", *m.CutoffMethod); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateDetectBladeCrash(formats strfmt.Registry) error {

	if err := validate.Required("detectBladeCrash", "body", m.DetectBladeCrash); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateDistortionAfterCutoffScaleFactor(formats strfmt.Registry) error {

	if swag.IsZero(m.DistortionAfterCutoffScaleFactor) { // not required
		return nil
	}

	if err := validate.Minimum("distortionAfterCutoffScaleFactor", "body", float64(*m.DistortionAfterCutoffScaleFactor), -5, false); err != nil {
		return err
	}

	if err := validate.Maximum("distortionAfterCutoffScaleFactor", "body", float64(*m.DistortionAfterCutoffScaleFactor), 5, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateDistortionAfterCutoffScaleFactorValues(formats strfmt.Registry) error {

	if swag.IsZero(m.DistortionAfterCutoffScaleFactorValues) { // not required
		return nil
	}

	for i := 0; i < len(m.DistortionAfterCutoffScaleFactorValues); i++ {

		if swag.IsZero(m.DistortionAfterCutoffScaleFactorValues[i]) { // not required
			continue
		}

		if err := validate.Minimum("distortionAfterCutoffScaleFactorValues"+"."+strconv.Itoa(i), "body", float64(*m.DistortionAfterCutoffScaleFactorValues[i]), -5, false); err != nil {
			return err
		}

		if err := validate.Maximum("distortionAfterCutoffScaleFactorValues"+"."+strconv.Itoa(i), "body", float64(*m.DistortionAfterCutoffScaleFactorValues[i]), 5, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartBasedSimulationParameters) validateDistortionScaleFactor(formats strfmt.Registry) error {

	if swag.IsZero(m.DistortionScaleFactor) { // not required
		return nil
	}

	if err := validate.Minimum("distortionScaleFactor", "body", float64(*m.DistortionScaleFactor), -5, false); err != nil {
		return err
	}

	if err := validate.Maximum("distortionScaleFactor", "body", float64(*m.DistortionScaleFactor), 5, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateDistortionScaleFactorValues(formats strfmt.Registry) error {

	if swag.IsZero(m.DistortionScaleFactorValues) { // not required
		return nil
	}

	for i := 0; i < len(m.DistortionScaleFactorValues); i++ {

		if swag.IsZero(m.DistortionScaleFactorValues[i]) { // not required
			continue
		}

		if err := validate.Minimum("distortionScaleFactorValues"+"."+strconv.Itoa(i), "body", float64(*m.DistortionScaleFactorValues[i]), -5, false); err != nil {
			return err
		}

		if err := validate.Maximum("distortionScaleFactorValues"+"."+strconv.Itoa(i), "body", float64(*m.DistortionScaleFactorValues[i]), 5, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartBasedSimulationParameters) validateElasticModulus(formats strfmt.Registry) error {

	if err := validate.Required("elasticModulus", "body", m.ElasticModulus); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateGenerateSupportVoxels(formats strfmt.Registry) error {

	if err := validate.Required("generateSupportVoxels", "body", m.GenerateSupportVoxels); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateHardeningFactor(formats strfmt.Registry) error {

	if err := validate.Required("hardeningFactor", "body", m.HardeningFactor); err != nil {
		return err
	}

	if err := validate.Minimum("hardeningFactor", "body", float64(*m.HardeningFactor), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("hardeningFactor", "body", float64(*m.HardeningFactor), 0.5, false); err != nil {
		return err
	}

	return nil
}

var partBasedSimulationParametersTypeLoadSteppingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DynamicLoadStepping","FixedLoadStepping"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partBasedSimulationParametersTypeLoadSteppingTypePropEnum = append(partBasedSimulationParametersTypeLoadSteppingTypePropEnum, v)
	}
}

const (
	// PartBasedSimulationParametersLoadSteppingTypeDynamicLoadStepping captures enum value "DynamicLoadStepping"
	PartBasedSimulationParametersLoadSteppingTypeDynamicLoadStepping string = "DynamicLoadStepping"
	// PartBasedSimulationParametersLoadSteppingTypeFixedLoadStepping captures enum value "FixedLoadStepping"
	PartBasedSimulationParametersLoadSteppingTypeFixedLoadStepping string = "FixedLoadStepping"
)

// prop value enum
func (m *PartBasedSimulationParameters) validateLoadSteppingTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partBasedSimulationParametersTypeLoadSteppingTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartBasedSimulationParameters) validateLoadSteppingType(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadSteppingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateLoadSteppingTypeEnum("loadSteppingType", "body", m.LoadSteppingType); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateLoadSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadSteps) { // not required
		return nil
	}

	if err := validate.MinimumInt("loadSteps", "body", int64(m.LoadSteps), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("loadSteps", "body", int64(m.LoadSteps), 200, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateMaximumThickWallThickness(formats strfmt.Registry) error {

	if swag.IsZero(m.MaximumThickWallThickness) { // not required
		return nil
	}

	if err := validate.Minimum("maximumThickWallThickness", "body", float64(m.MaximumThickWallThickness), 2e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("maximumThickWallThickness", "body", float64(m.MaximumThickWallThickness), 0.01, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateMaximumThinWallDistance(formats strfmt.Registry) error {

	if swag.IsZero(m.MaximumThinWallDistance) { // not required
		return nil
	}

	if err := validate.Minimum("maximumThinWallDistance", "body", float64(m.MaximumThinWallDistance), 2e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("maximumThinWallDistance", "body", float64(m.MaximumThinWallDistance), 0.01, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateMinimumSupportHeight(formats strfmt.Registry) error {

	if err := validate.Required("minimumSupportHeight", "body", m.MinimumSupportHeight); err != nil {
		return err
	}

	if err := validate.Minimum("minimumSupportHeight", "body", float64(*m.MinimumSupportHeight), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("minimumSupportHeight", "body", float64(*m.MinimumSupportHeight), 0.005, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateMinimumThickWallThickness(formats strfmt.Registry) error {

	if swag.IsZero(m.MinimumThickWallThickness) { // not required
		return nil
	}

	if err := validate.Minimum("minimumThickWallThickness", "body", float64(m.MinimumThickWallThickness), 2e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("minimumThickWallThickness", "body", float64(m.MinimumThickWallThickness), 0.01, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateOutputDisplacementAfterCutoff(formats strfmt.Registry) error {

	if err := validate.Required("outputDisplacementAfterCutoff", "body", m.OutputDisplacementAfterCutoff); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateOutputLayerVtk(formats strfmt.Registry) error {

	if err := validate.Required("outputLayerVtk", "body", m.OutputLayerVtk); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateOutputMapdl(formats strfmt.Registry) error {

	if err := validate.Required("outputMapdl", "body", m.OutputMapdl); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validatePerformSupportOnlyCutoff(formats strfmt.Registry) error {

	if err := validate.Required("performSupportOnlyCutoff", "body", m.PerformSupportOnlyCutoff); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validatePerformSupportOptimization(formats strfmt.Registry) error {

	if err := validate.Required("performSupportOptimization", "body", m.PerformSupportOptimization); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validatePoissonRatio(formats strfmt.Registry) error {

	if err := validate.Required("poissonRatio", "body", m.PoissonRatio); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSimulationParts(formats strfmt.Registry) error {

	if swag.IsZero(m.SimulationParts) { // not required
		return nil
	}

	for i := 0; i < len(m.SimulationParts); i++ {

		if swag.IsZero(m.SimulationParts[i]) { // not required
			continue
		}

		if m.SimulationParts[i] != nil {

			if err := m.SimulationParts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("simulationParts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PartBasedSimulationParameters) validateStrainScalingFactor(formats strfmt.Registry) error {

	if err := validate.Required("strainScalingFactor", "body", m.StrainScalingFactor); err != nil {
		return err
	}

	return nil
}

var partBasedSimulationParametersTypeStressModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LinearElastic","J2Plasticity"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partBasedSimulationParametersTypeStressModePropEnum = append(partBasedSimulationParametersTypeStressModePropEnum, v)
	}
}

const (
	// PartBasedSimulationParametersStressModeLinearElastic captures enum value "LinearElastic"
	PartBasedSimulationParametersStressModeLinearElastic string = "LinearElastic"
	// PartBasedSimulationParametersStressModeJ2Plasticity captures enum value "J2Plasticity"
	PartBasedSimulationParametersStressModeJ2Plasticity string = "J2Plasticity"
)

// prop value enum
func (m *PartBasedSimulationParameters) validateStressModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partBasedSimulationParametersTypeStressModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartBasedSimulationParameters) validateStressMode(formats strfmt.Registry) error {

	if err := validate.Required("stressMode", "body", m.StressMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateStressModeEnum("stressMode", "body", *m.StressMode); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSupportAngle(formats strfmt.Registry) error {

	if err := validate.Required("supportAngle", "body", m.SupportAngle); err != nil {
		return err
	}

	if err := validate.Minimum("supportAngle", "body", float64(*m.SupportAngle), 1, false); err != nil {
		return err
	}

	if err := validate.Maximum("supportAngle", "body", float64(*m.SupportAngle), 89, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSupportFactorOfSafety(formats strfmt.Registry) error {

	if err := validate.Required("supportFactorOfSafety", "body", m.SupportFactorOfSafety); err != nil {
		return err
	}

	if err := validate.Minimum("supportFactorOfSafety", "body", float64(*m.SupportFactorOfSafety), 0.1, false); err != nil {
		return err
	}

	if err := validate.Maximum("supportFactorOfSafety", "body", float64(*m.SupportFactorOfSafety), 10, false); err != nil {
		return err
	}

	return nil
}

var partBasedSimulationParametersTypeSupportTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Generated","UserDefined"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partBasedSimulationParametersTypeSupportTypePropEnum = append(partBasedSimulationParametersTypeSupportTypePropEnum, v)
	}
}

const (
	// PartBasedSimulationParametersSupportTypeGenerated captures enum value "Generated"
	PartBasedSimulationParametersSupportTypeGenerated string = "Generated"
	// PartBasedSimulationParametersSupportTypeUserDefined captures enum value "UserDefined"
	PartBasedSimulationParametersSupportTypeUserDefined string = "UserDefined"
)

// prop value enum
func (m *PartBasedSimulationParameters) validateSupportTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partBasedSimulationParametersTypeSupportTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartBasedSimulationParameters) validateSupportType(formats strfmt.Registry) error {

	if err := validate.Required("supportType", "body", m.SupportType); err != nil {
		return err
	}

	// value enum
	if err := m.validateSupportTypeEnum("supportType", "body", *m.SupportType); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSupportYieldStrength(formats strfmt.Registry) error {

	if err := validate.Required("supportYieldStrength", "body", m.SupportYieldStrength); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSupportYieldStrengthRatio(formats strfmt.Registry) error {

	if err := validate.Required("supportYieldStrengthRatio", "body", m.SupportYieldStrengthRatio); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateThickWallDistance(formats strfmt.Registry) error {

	if swag.IsZero(m.ThickWallDistance) { // not required
		return nil
	}

	if err := validate.Minimum("thickWallDistance", "body", float64(m.ThickWallDistance), 2e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("thickWallDistance", "body", float64(m.ThickWallDistance), 0.01, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateThinWallThickness(formats strfmt.Registry) error {

	if swag.IsZero(m.ThinWallThickness) { // not required
		return nil
	}

	if err := validate.Minimum("thinWallThickness", "body", float64(m.ThinWallThickness), 2e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("thinWallThickness", "body", float64(m.ThinWallThickness), 0.0005, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateVoxelSampleRate(formats strfmt.Registry) error {

	if swag.IsZero(m.VoxelSampleRate) { // not required
		return nil
	}

	if err := validate.MinimumInt("voxelSampleRate", "body", int64(m.VoxelSampleRate), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("voxelSampleRate", "body", int64(m.VoxelSampleRate), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateVoxelSize(formats strfmt.Registry) error {

	if err := validate.Required("voxelSize", "body", m.VoxelSize); err != nil {
		return err
	}

	if err := validate.Minimum("voxelSize", "body", float64(*m.VoxelSize), 2e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("voxelSize", "body", float64(*m.VoxelSize), 0.002, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartBasedSimulationParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartBasedSimulationParameters) UnmarshalBinary(b []byte) error {
	var res PartBasedSimulationParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
