// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ThermalSimulationParameters thermal simulation parameters
// swagger:model ThermalSimulationParameters
type ThermalSimulationParameters struct {

	// anisotropic strain coefficients parallel
	// Required: true
	AnisotropicStrainCoefficientsParallel *float64 `json:"anisotropicStrainCoefficientsParallel"`

	// anisotropic strain coefficients perpendicular
	// Required: true
	AnisotropicStrainCoefficientsPerpendicular *float64 `json:"anisotropicStrainCoefficientsPerpendicular"`

	// anisotropic strain coefficients z
	// Required: true
	AnisotropicStrainCoefficientsZ *float64 `json:"anisotropicStrainCoefficientsZ"`

	// diameter of laser beam in meters
	// Required: true
	// Maximum: 0.00014
	// Minimum: 1e-05
	BeamDiameter *float64 `json:"beamDiameter"`

	// should be a number between 0.05 and 1.5 mm
	CoaxialAverageSensorRadius float64 `json:"coaxialAverageSensorRadius,omitempty"`

	// list of z height (mm) ranges within the part where the coaxial average sensor will collect data
	CoaxialAverageSensorZHeights []*ZHeightRange `json:"coaxialAverageSensorZHeights"`

	// Must be between 0.00001 to 0.001 meters
	// Required: true
	// Maximum: 0.001
	// Minimum: 1e-05
	HatchSpacing *float64 `json:"hatchSpacing"`

	// heater temperature in degrees kelvin
	// Maximum: 773.15
	// Minimum: 293.15
	HeaterTemperature float64 `json:"heaterTemperature,omitempty"`

	// false indicates that only the thermal solver will run, while true indicates that the mechanics solver will run after the thermal solver
	// Required: true
	IncludeStressAnalysis *bool `json:"includeStressAnalysis"`

	// Array of integer layer numbers where instant dynamic sensor data will be collected
	InstantDynamicSensorLayers []int32 `json:"instantDynamicSensorLayers"`

	// radius for instant dynamic sensor data collection in mm
	// Maximum: 1.5
	// Minimum: 0.05
	InstantDynamicSensorRadius float64 `json:"instantDynamicSensorRadius,omitempty"`

	// radius for instant static sensor data collection in mm
	// Maximum: 1.5
	// Minimum: 0.05
	InstantStaticSensorRadius float64 `json:"instantStaticSensorRadius,omitempty"`

	// Must be between 50 to 700 watts
	// Required: true
	// Maximum: 700
	// Minimum: 50
	LaserWattage *float64 `json:"laserWattage"`

	// Must be between 0 to 180 degrees
	// Required: true
	// Maximum: 180
	// Minimum: 0
	LayerRotationAngle *float64 `json:"layerRotationAngle"`

	// Must be between 0.00001 to 0.0001 meters
	// Required: true
	// Maximum: 0.0001
	// Minimum: 1e-05
	LayerThickness *float64 `json:"layerThickness"`

	// multiplier of fine grid spacing for course grid portion of FEA mesh
	// Maximum: 12
	// Minimum: 1
	MeshResolutionFactor int32 `json:"meshResolutionFactor,omitempty"`

	// if true, coaxial average sensor data will be collected for the deposity layers specified by coaxialAverageSensorZHeights
	// Required: true
	OutputCoaxialAverageSensorData *bool `json:"outputCoaxialAverageSensorData"`

	// if true, dyanmic sensor data will be collected for each layer specified in the instantDynamicSensorLayers property
	// Required: true
	OutputInstantDynamicSensor *bool `json:"outputInstantDynamicSensor"`

	// if true, instant static sensor data will be collected for each selectedPoint property
	// Required: true
	OutputInstantStaticSensor *bool `json:"outputInstantStaticSensor"`

	// if true, probe sensor data will be collected for each selectedPoint property
	// Required: true
	OutputPointProbe *bool `json:"outputPointProbe"`

	// for each slectedPoint, a series of vtk files will output thermal history around that point with a radius of staticVirtualSensorRadius.
	// Required: true
	OutputPointThermalHistory *bool `json:"outputPointThermalHistory"`

	// if true, PrintRite sensor data will be collected for each selectedPoint property
	// Required: true
	OutputPrintRitePcsSensor *bool `json:"outputPrintRitePcsSensor"`

	// output shrinkage
	// Required: true
	OutputShrinkage *bool `json:"outputShrinkage"`

	// output state map
	// Required: true
	OutputStateMap *bool `json:"outputStateMap"`

	// output thermal vtk
	OutputThermalVtk bool `json:"outputThermalVtk,omitempty"`

	// output thermal vtk layers
	OutputThermalVtkLayers string `json:"outputThermalVtkLayers,omitempty"`

	// radius for PrintRite sensor data collection in mm
	// Maximum: 1.5
	// Minimum: 0.05
	PrintRitePcsSensorRadius float64 `json:"printRitePcsSensorRadius,omitempty"`

	// if true, pyrometer sensor data will be collected for every layer
	PyroVirtualSensorOutputAllLayers bool `json:"pyroVirtualSensorOutputAllLayers,omitempty"`

	// Must be between 0.35 to 2.5 meters/second
	// Required: true
	// Maximum: 2.5
	// Minimum: 0.35
	ScanSpeed *float64 `json:"scanSpeed"`

	// List of points where the thermal solver will collect thermal history - limit 10
	SelectedPoints []*SelectedPoint `json:"selectedPoints"`

	// Must be between 0.001 to 0.1 meters
	// Required: true
	// Maximum: 0.1
	// Minimum: 0.001
	SlicingStripeWidth *float64 `json:"slicingStripeWidth"`

	// Must be between 0 to 180 degrees
	// Required: true
	// Maximum: 180
	// Minimum: 0
	StartingLayerAngle *float64 `json:"startingLayerAngle"`

	// should be a number between 0.05 and 5.0 mm
	StaticVirtualSensorRadius float64 `json:"staticVirtualSensorRadius,omitempty"`
}

// Validate validates this thermal simulation parameters
func (m *ThermalSimulationParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnisotropicStrainCoefficientsParallel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAnisotropicStrainCoefficientsPerpendicular(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAnisotropicStrainCoefficientsZ(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBeamDiameter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCoaxialAverageSensorZHeights(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHatchSpacing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHeaterTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIncludeStressAnalysis(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstantDynamicSensorLayers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstantDynamicSensorRadius(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstantStaticSensorRadius(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLaserWattage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerRotationAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerThickness(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeshResolutionFactor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputCoaxialAverageSensorData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputInstantDynamicSensor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputInstantStaticSensor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputPointProbe(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputPointThermalHistory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputPrintRitePcsSensor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputShrinkage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputStateMap(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrintRitePcsSensorRadius(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScanSpeed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSelectedPoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSlicingStripeWidth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartingLayerAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThermalSimulationParameters) validateAnisotropicStrainCoefficientsParallel(formats strfmt.Registry) error {

	if err := validate.Required("anisotropicStrainCoefficientsParallel", "body", m.AnisotropicStrainCoefficientsParallel); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateAnisotropicStrainCoefficientsPerpendicular(formats strfmt.Registry) error {

	if err := validate.Required("anisotropicStrainCoefficientsPerpendicular", "body", m.AnisotropicStrainCoefficientsPerpendicular); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateAnisotropicStrainCoefficientsZ(formats strfmt.Registry) error {

	if err := validate.Required("anisotropicStrainCoefficientsZ", "body", m.AnisotropicStrainCoefficientsZ); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateBeamDiameter(formats strfmt.Registry) error {

	if err := validate.Required("beamDiameter", "body", m.BeamDiameter); err != nil {
		return err
	}

	if err := validate.Minimum("beamDiameter", "body", float64(*m.BeamDiameter), 1e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("beamDiameter", "body", float64(*m.BeamDiameter), 0.00014, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateCoaxialAverageSensorZHeights(formats strfmt.Registry) error {

	if swag.IsZero(m.CoaxialAverageSensorZHeights) { // not required
		return nil
	}

	for i := 0; i < len(m.CoaxialAverageSensorZHeights); i++ {

		if swag.IsZero(m.CoaxialAverageSensorZHeights[i]) { // not required
			continue
		}

		if m.CoaxialAverageSensorZHeights[i] != nil {

			if err := m.CoaxialAverageSensorZHeights[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coaxialAverageSensorZHeights" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThermalSimulationParameters) validateHatchSpacing(formats strfmt.Registry) error {

	if err := validate.Required("hatchSpacing", "body", m.HatchSpacing); err != nil {
		return err
	}

	if err := validate.Minimum("hatchSpacing", "body", float64(*m.HatchSpacing), 1e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("hatchSpacing", "body", float64(*m.HatchSpacing), 0.001, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateHeaterTemperature(formats strfmt.Registry) error {

	if swag.IsZero(m.HeaterTemperature) { // not required
		return nil
	}

	if err := validate.Minimum("heaterTemperature", "body", float64(m.HeaterTemperature), 293.15, false); err != nil {
		return err
	}

	if err := validate.Maximum("heaterTemperature", "body", float64(m.HeaterTemperature), 773.15, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateIncludeStressAnalysis(formats strfmt.Registry) error {

	if err := validate.Required("includeStressAnalysis", "body", m.IncludeStressAnalysis); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateInstantDynamicSensorLayers(formats strfmt.Registry) error {

	if swag.IsZero(m.InstantDynamicSensorLayers) { // not required
		return nil
	}

	return nil
}

func (m *ThermalSimulationParameters) validateInstantDynamicSensorRadius(formats strfmt.Registry) error {

	if swag.IsZero(m.InstantDynamicSensorRadius) { // not required
		return nil
	}

	if err := validate.Minimum("instantDynamicSensorRadius", "body", float64(m.InstantDynamicSensorRadius), 0.05, false); err != nil {
		return err
	}

	if err := validate.Maximum("instantDynamicSensorRadius", "body", float64(m.InstantDynamicSensorRadius), 1.5, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateInstantStaticSensorRadius(formats strfmt.Registry) error {

	if swag.IsZero(m.InstantStaticSensorRadius) { // not required
		return nil
	}

	if err := validate.Minimum("instantStaticSensorRadius", "body", float64(m.InstantStaticSensorRadius), 0.05, false); err != nil {
		return err
	}

	if err := validate.Maximum("instantStaticSensorRadius", "body", float64(m.InstantStaticSensorRadius), 1.5, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateLaserWattage(formats strfmt.Registry) error {

	if err := validate.Required("laserWattage", "body", m.LaserWattage); err != nil {
		return err
	}

	if err := validate.Minimum("laserWattage", "body", float64(*m.LaserWattage), 50, false); err != nil {
		return err
	}

	if err := validate.Maximum("laserWattage", "body", float64(*m.LaserWattage), 700, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateLayerRotationAngle(formats strfmt.Registry) error {

	if err := validate.Required("layerRotationAngle", "body", m.LayerRotationAngle); err != nil {
		return err
	}

	if err := validate.Minimum("layerRotationAngle", "body", float64(*m.LayerRotationAngle), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("layerRotationAngle", "body", float64(*m.LayerRotationAngle), 180, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateLayerThickness(formats strfmt.Registry) error {

	if err := validate.Required("layerThickness", "body", m.LayerThickness); err != nil {
		return err
	}

	if err := validate.Minimum("layerThickness", "body", float64(*m.LayerThickness), 1e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("layerThickness", "body", float64(*m.LayerThickness), 0.0001, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateMeshResolutionFactor(formats strfmt.Registry) error {

	if swag.IsZero(m.MeshResolutionFactor) { // not required
		return nil
	}

	if err := validate.MinimumInt("meshResolutionFactor", "body", int64(m.MeshResolutionFactor), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("meshResolutionFactor", "body", int64(m.MeshResolutionFactor), 12, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputCoaxialAverageSensorData(formats strfmt.Registry) error {

	if err := validate.Required("outputCoaxialAverageSensorData", "body", m.OutputCoaxialAverageSensorData); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputInstantDynamicSensor(formats strfmt.Registry) error {

	if err := validate.Required("outputInstantDynamicSensor", "body", m.OutputInstantDynamicSensor); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputInstantStaticSensor(formats strfmt.Registry) error {

	if err := validate.Required("outputInstantStaticSensor", "body", m.OutputInstantStaticSensor); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputPointProbe(formats strfmt.Registry) error {

	if err := validate.Required("outputPointProbe", "body", m.OutputPointProbe); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputPointThermalHistory(formats strfmt.Registry) error {

	if err := validate.Required("outputPointThermalHistory", "body", m.OutputPointThermalHistory); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputPrintRitePcsSensor(formats strfmt.Registry) error {

	if err := validate.Required("outputPrintRitePcsSensor", "body", m.OutputPrintRitePcsSensor); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputShrinkage(formats strfmt.Registry) error {

	if err := validate.Required("outputShrinkage", "body", m.OutputShrinkage); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputStateMap(formats strfmt.Registry) error {

	if err := validate.Required("outputStateMap", "body", m.OutputStateMap); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validatePrintRitePcsSensorRadius(formats strfmt.Registry) error {

	if swag.IsZero(m.PrintRitePcsSensorRadius) { // not required
		return nil
	}

	if err := validate.Minimum("printRitePcsSensorRadius", "body", float64(m.PrintRitePcsSensorRadius), 0.05, false); err != nil {
		return err
	}

	if err := validate.Maximum("printRitePcsSensorRadius", "body", float64(m.PrintRitePcsSensorRadius), 1.5, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateScanSpeed(formats strfmt.Registry) error {

	if err := validate.Required("scanSpeed", "body", m.ScanSpeed); err != nil {
		return err
	}

	if err := validate.Minimum("scanSpeed", "body", float64(*m.ScanSpeed), 0.35, false); err != nil {
		return err
	}

	if err := validate.Maximum("scanSpeed", "body", float64(*m.ScanSpeed), 2.5, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateSelectedPoints(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectedPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.SelectedPoints); i++ {

		if swag.IsZero(m.SelectedPoints[i]) { // not required
			continue
		}

		if m.SelectedPoints[i] != nil {

			if err := m.SelectedPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectedPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThermalSimulationParameters) validateSlicingStripeWidth(formats strfmt.Registry) error {

	if err := validate.Required("slicingStripeWidth", "body", m.SlicingStripeWidth); err != nil {
		return err
	}

	if err := validate.Minimum("slicingStripeWidth", "body", float64(*m.SlicingStripeWidth), 0.001, false); err != nil {
		return err
	}

	if err := validate.Maximum("slicingStripeWidth", "body", float64(*m.SlicingStripeWidth), 0.1, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateStartingLayerAngle(formats strfmt.Registry) error {

	if err := validate.Required("startingLayerAngle", "body", m.StartingLayerAngle); err != nil {
		return err
	}

	if err := validate.Minimum("startingLayerAngle", "body", float64(*m.StartingLayerAngle), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("startingLayerAngle", "body", float64(*m.StartingLayerAngle), 180, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThermalSimulationParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThermalSimulationParameters) UnmarshalBinary(b []byte) error {
	var res ThermalSimulationParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
