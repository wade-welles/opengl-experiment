package opengl

import (
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
)

// GLDataTypes
var (
	GLDataTypeFloat = GLDataType{
		name:  "FLOAT",
		size:  4,
		value: gl.FLOAT,
	}
	GLDataTypeInt = GLDataType{
		name:  "INT",
		size:  4,
		value: gl.INT,
	}
	GLDataTypeUint = GLDataType{
		name:  "UNSIGNED_INT",
		size:  4,
		value: gl.UNSIGNED_INT,
	}
)

// GLDataType .
type GLDataType struct {
	name  string
	size  int
	value uint32
}

// VBO .
type VBO struct {
	id     uint32
	layout *VBOLayout
}

// NewVBO .
func NewVBO(size int) (*VBO, error) {
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	vbo := &VBO{id: vboID}

	vbo.Bind()
	defer vbo.Unbind()

	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(nil), gl.DYNAMIC_DRAW)

	return vbo, nil
}

// SetData .
func (v *VBO) SetData(data VBOData) {
	v.Bind()
	defer v.Unbind()

	gl.BufferSubData(gl.ARRAY_BUFFER, 0, data.GetVBOSize(), data.GetVBOGLPtr())
}

// GetLayout .
func (v *VBO) GetLayout() *VBOLayout {
	return v.layout
}

// SetLayout .
func (v *VBO) SetLayout(layout *VBOLayout) {
	v.layout = layout
}

// Bind .
func (v *VBO) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, v.id)
}

// Unbind .
func (v *VBO) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// VBOLayout .
type VBOLayout struct {
	stride   int32
	elements []VBOLayoutElement
}

// NewVBOLayout .
func NewVBOLayout(elements ...VBOLayoutElement) *VBOLayout {
	layout := &VBOLayout{}
	for _, e := range elements {
		layout.elements = append(layout.elements, e)
		layout.stride += (int32(e.DataType.size) * e.Count)
	}
	return layout
}

// GetStride .
func (l *VBOLayout) GetStride() int32 {
	return l.stride
}

// VBOLayoutElement .
type VBOLayoutElement struct {
	Count      int32
	Normalized bool
	DataType   GLDataType
}

// VBOData .
type VBOData interface {
	GetVBOGLPtr() unsafe.Pointer
	GetVBOSize() int
	GetVertexSize() int
}
