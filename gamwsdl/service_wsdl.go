package gamwsdl

import "encoding/xml"

type WSDLDefinition struct {
	XMLName         xml.Name `xml:"definitions"`
	TargetNamespace string   `xml:"targetNamespace,attr"`
	Service         Service  `xml:"service"`
	Types           *Types   `xml:"types"`
	PortType        PortType `xml:"portType"`
	Binding         Binding  `xml:"binding"`
}

type Service struct {
	XMLName xml.Name `xml:"service"`
	Name    string   `xml:"name,attr"`
	Port    Port     `xml:"port"`
}
type Types struct {
	XMLName xml.Name `xml:"types"`
	Schema  *Schema  `xml:"schema"`
}

type Schema struct {
	XMLName      xml.Name          `xml:"schema"`
	Xmlns        map[string]string `xml:"-"`
	TNS          string            `xml:"xmlns tns,attr"`
	XSD          string            `xml:"xmlns xs,attr"`
	ComplexTypes []ComplexType     `xml:"complexType"`
	SimpleTypes  []SimpleType      `xml:"simpleType"`
	Elements     []Element         `xml:"element"`
}

type ComplexType struct {
	Name string `xml:"name,attr"`
	// abstract="true" should be implemented as interface
	// an interface has some methods
	//
	Abstract       *bool           `xml:"abstract,attr"`
	Annotation     *Annotation     `xml:"annotation"`
	Sequence       *Sequence       `xml:"sequence"`
	ComplexContent *ComplexContent `xml:"complexContent"`
}

type ComplexContent struct {
	// Extensionは何かのクラスを継承している
	Extension *Extension `xml:"extension"`
}

type SimpleType struct {
	Name        string       `xml:"name,attr"`
	Annotation  *Annotation  `xml:"annotation"`
	Restriction *Restriction `xml:"restriction"`
}

type Annotation struct {
	Documentation *string `xml:"documentation"`
}

type Extension struct {
	Base     string    `xml:"base,attr"`
	Sequence *Sequence `xml:"sequence"`
}

type Sequence struct {
	Elements []Element `xml:"element"`
}

type Element struct {
	// 1, unbounded
	MaxOccurs *string `xml:"maxOccurs,attr"`
	// 0
	MinOccurs *string `xml:"minOccurs,attr"`
	Name      string  `xml:"name,attr"`
	// ex: type="xsd:string"
	Type        *string      `xml:"type,attr"`
	Annotation  *Annotation  `xml:"annotation"`
	ComplexType *ComplexType `xml:"complexType"`
}

type Restriction struct {
	Base         string        `xml:"base,attr"`
	Enumerations []Enumeration `xml:"enumeration"`
}

type Enumeration struct {
	Value      string      `xml:"value,attr"`
	Annotation *Annotation `xml:"annotation"`
}

type PortType struct {
	Documentations *string     `xml:"documentation"`
	Name           string      `xml:"name,attr"`
	Operations     []Operation `xml:"operation"`
}

type Binding struct {
	XMLName    xml.Name    `xml:"binding"`
	Name       string      `xml:"name,attr"`
	Operations []Operation `xml:"operation"`
}

type Operation struct {
	XMLName        xml.Name `xml:"operation"`
	Documentations *string  `xml:"documentation"`
	Name           string   `xml:"name,attr"`
	Input          Input    `xml:"input"`
	Output         Output   `xml:"output"`
}

type Input struct {
	Name    string `xml:"name,attr"`
	Message string `xml:"message,attr"`
}

type Output struct {
	Name string `xml:"name,attr"`
}

type Port struct {
	Name    string  `xml:"name,attr"`
	Address Address `xml:"address"`
}

type Address struct {
	Location string `xml:"location,attr"`
}
