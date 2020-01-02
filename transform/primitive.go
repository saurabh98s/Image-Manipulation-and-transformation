package primitive

type Mode int
const (

	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)
// withMode is an option for the trnsform function that will define
// the mode you want to use.Default= Modetriangle
func WithMode(mode Mode) func() []string {
	return func()string {
		return []string{"-n", fmt.Sprintf("%d",mode)}
	}
}
// will apply a primitive transformation and return a resulting
// images to the reader
func Transform(image io.Reader,numShapes int,opts ...func() []string) io.Reader{
	in,err:=ioutil.TempFile("","in_")
	if err != nil {
		return nil, err
	}
	// a better option would be to handle the error insted of
	// just returning it.
	defer os.Remove(in.Name())
	// removes the created temp file in case there is an error.
	out,err:=ioutil.TempFile("","out_")
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())
// read image into input file
_,err= io.Copy(in,image)
if err != nil {
	return nil, err
}

// Run primitive w/ -i in.Name() -o out.Name()
stdCombo,err:=primitive(in.Name(),out.Name(),numShapes,ModeCombo)
if err != nil {
	return nil, err
}
fmt.Println(stdCombo)
// read out into a reader, return reader, delete out
b:=bytes.NewBuffer(nil)
_,err:=io.Copy(b,out)
if err != nil {
	return nil, err
}
return b,nil

} 

func primitive(inputFile,outputFile string,numShapes int, modes PrimitiveMode)(string,error){
	// https://github.com/fogleman/primitive
	// the exec package takes two parameters the  command (here "primitive") and  the parameters
	argStr := fmt.Sprintf("-i %s -o %s -n %d -m %d",inputFile,outputFile,numShapes,modes)
	cmd:=exec.Command("primitive", strings.Fields(argStr)...)
	b,err:=cmd.CombinedOutput()
	return string(b),err
	


}
func tempfile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", "in_")
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}