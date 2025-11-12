// Interface Segregation Principle (ISP)
// Shouldnt push too much in the same interface

package main

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {
	panic("implement me")
}

func (m MultiFunctionPrinter) Fax(d Document) {
	panic("implement me")
}

func (m MultiFunctionPrinter) Scan(d Document) {
	panic("implement me")
}

type OldFashionedPrinter struct{}

func (o OldFashionedPrinter) Print(d Document) {
	panic("implement me")
}

// Deprecated: ...
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// IS -> Try to break the interface in functions that the people would need
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Faxer interface {
	Fax(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {

}

type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	panic("implement me")
}

func (p Photocopier) Print(d Document) {
	panic("implement me")
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// decorator design pattern
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}
func main() {

}
