type {{$.Name}}GinSrv struct {

}

func (s *{{$.Name}}GinSrv) Say{{$.Name}} (req *Say{{$.Name}}Req) (*Say{{$.Name}}Rep, error) {

	return &Say{{$.Name}}Rep{Reply: req.Name}, nil
}