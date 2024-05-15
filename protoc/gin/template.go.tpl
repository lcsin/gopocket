type {{$.Name}}GinSrv struct {
	s HelloServer
}

{{range .Methods}}
{{.Comment}}func (h *{{$.Name}}GinSrv) {{.Name}} (c *gin.Context) {
	var req {{.Request}}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	resp, err := h.s.{{.Name}}(c, &{{.Request}}{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp.Reply)
}
{{end}}