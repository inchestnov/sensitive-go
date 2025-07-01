package sensitive

import (
	"reflect"
	"testing"
)

func TestDetachStringSensitiveField(t *testing.T) {
	t.Parallel()

	type User struct {
		Password string `sensitive:"true"`
	}
	u := User{Password: "Password"}

	insensitive, _, err := Detach(u)
	assertNoError(t, err)
	assertTrue(t, insensitive.Password == "", "string sensitive field should be empty")
}

func TestDetachOnlySensitiveField(t *testing.T) {
	t.Parallel()

	type User struct {
		Username string
		Password string `sensitive:"true"`
	}
	u := User{Username: "username", Password: "Password"}

	insensitive, _, err := Detach(u)
	assertNoError(t, err)
	assertTrue(t, insensitive.Username == u.Username, "non-sensitive fields should not changed")
}

func TestDetachOnlySensitiveTrueField(t *testing.T) {
	t.Parallel()

	type User struct {
		Username string `sensitive:"false"`
		Password string `sensitive:"true"`
	}
	u := User{Username: "username", Password: "Password"}

	insensitive, _, err := Detach(u)
	assertNoError(t, err)
	assertTrue(t, insensitive.Username == u.Username, "non-sensitive fields should not changed")
}

func TestDetachPointerField(t *testing.T) {
	t.Parallel()

	type User struct {
		Password *string `sensitive:"true"`
	}
	pass := "Password"
	u := User{Password: &pass}

	insensitive, _, err := Detach(u)
	assertNoError(t, err)
	assertTrue(t, insensitive.Password == nil, "pointer sensitive field should be nil")
}

func TestDetachNumbers(t *testing.T) {
	t.Parallel()

	type Numbers struct {
		Int   int   `sensitive:"true"`
		Int8  int8  `sensitive:"true"`
		Int16 int16 `sensitive:"true"`
		Int32 int32 `sensitive:"true"`
		Int64 int64 `sensitive:"true"`

		UInt   uint   `sensitive:"true"`
		UInt8  uint8  `sensitive:"true"`
		UInt16 uint16 `sensitive:"true"`
		UInt32 uint32 `sensitive:"true"`
		UInt64 uint64 `sensitive:"true"`

		Float32 float32 `sensitive:"true"`
		Float64 float64 `sensitive:"true"`

		Complex64  complex64  `sensitive:"true"`
		Complex128 complex128 `sensitive:"true"`
	}

	tests := []struct {
		name   string
		source Numbers
		want   Numbers
	}{
		{
			name:   "int",
			source: Numbers{Int: 1},
			want:   Numbers{Int: 0},
		},
		{
			name:   "int8",
			source: Numbers{Int8: 1},
			want:   Numbers{Int8: 0},
		},
		{
			name:   "int16",
			source: Numbers{Int16: 1},
			want:   Numbers{Int16: 0},
		},
		{
			name:   "int32",
			source: Numbers{Int32: 1},
			want:   Numbers{Int32: 0},
		},
		{
			name:   "int64",
			source: Numbers{Int64: 1},
			want:   Numbers{Int64: 0},
		},

		{
			name:   "uint",
			source: Numbers{Int: 1},
			want:   Numbers{Int: 0},
		},
		{
			name:   "uint8",
			source: Numbers{UInt8: 1},
			want:   Numbers{UInt8: 0},
		},
		{
			name:   "uint16",
			source: Numbers{UInt16: 1},
			want:   Numbers{UInt16: 0},
		},
		{
			name:   "uint32",
			source: Numbers{UInt32: 1},
			want:   Numbers{UInt32: 0},
		},
		{
			name:   "uint64",
			source: Numbers{UInt64: 1},
			want:   Numbers{UInt64: 0},
		},

		{
			name:   "float32",
			source: Numbers{Float32: 1.0},
			want:   Numbers{Float32: 0},
		},
		{
			name:   "float64",
			source: Numbers{Float64: 1.0},
			want:   Numbers{Float64: 0},
		},

		{
			name:   "complex64",
			source: Numbers{Complex64: complex(1, 1)},
			want:   Numbers{Complex64: complex(0, 0)},
		},
		{
			name:   "complex64 only real",
			source: Numbers{Complex64: complex(1, 0)},
			want:   Numbers{Complex64: complex(0, 0)},
		},
		{
			name:   "complex64 only imagine",
			source: Numbers{Complex64: complex(0, 1)},
			want:   Numbers{Complex64: complex(0, 0)},
		},
		{
			name:   "complex128",
			source: Numbers{Complex128: complex(1, 2)},
			want:   Numbers{Complex128: complex(0, 0)},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			insensitive, _, err := Detach(tt.source)
			assertNoError(t, err)
			if !reflect.DeepEqual(insensitive, tt.want) {
				t.Errorf("Detach() = %v, want %v", insensitive, tt.want)
			}
		})
	}
}

func TestDetachSlices(t *testing.T) {
	t.Parallel()

	type Slices struct {
		Slice []int `sensitive:"true"`
	}

	tests := []struct {
		name   string
		source Slices
		want   Slices
	}{
		{
			name:   "empty slice -> nil",
			source: Slices{Slice: []int{}},
			want:   Slices{Slice: nil},
		},
		{
			name:   "one-item slice -> nil",
			source: Slices{Slice: []int{1}},
			want:   Slices{Slice: nil},
		},
		{
			name:   "zero-item slice -> nil",
			source: Slices{Slice: []int{0}},
			want:   Slices{Slice: nil},
		},
		{
			name:   "10-item slice -> nil",
			source: Slices{Slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want:   Slices{Slice: nil},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			insensitive, _, err := Detach(tt.source)
			assertNoError(t, err)
			if !reflect.DeepEqual(insensitive, tt.want) {
				t.Errorf("Detach() = %v, want %v", insensitive, tt.want)
			}
		})
	}
}

func TestDetachMaps(t *testing.T) {
	type Maps struct {
		Map map[string]string `sensitive:"true"`
	}

	tests := []struct {
		name   string
		source Maps
		want   Maps
	}{
		{
			name:   "empty map -> nil",
			source: Maps{Map: map[string]string{}},
			want:   Maps{Map: nil},
		},
		{
			name:   "one-key map -> nil",
			source: Maps{Map: map[string]string{"key": "value"}},
			want:   Maps{Map: nil},
		},
		{
			name:   "zero-item map -> nil",
			source: Maps{Map: map[string]string{"": ""}},
			want:   Maps{Map: nil},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			insensitive, _, err := Detach(tt.source)
			assertNoError(t, err)
			if !reflect.DeepEqual(insensitive, tt.want) {
				t.Errorf("Detach() = %v, want %v", insensitive, tt.want)
			}
		})
	}
}

func TestDetachUnexported(t *testing.T) {
	t.Parallel()

	type PrivateFieldStruct struct {
		field string `sensitive:"true"`
	}
	s := PrivateFieldStruct{field: "value"}

	insensitive, _, err := Detach(s)
	assertNoError(t, err)
	assertTrue(t, insensitive.field == s.field, "unexported sensitive field should not changed")
}

func TestDetachInterface(t *testing.T) {
	t.Parallel()

	type Interface interface{}
	type Struct struct {
		InterfaceField Interface `sensitive:"true"`
	}

	s := Struct{InterfaceField: "value"}

	detached, _, err := Detach(s)
	assertNoError(t, err)

	assertTrue(t, detached.InterfaceField == nil, "interface sensitive field should be null")
}

func TestDetachPointer(t *testing.T) {
	t.Parallel()

	type Struct struct {
		Field string `sensitive:"true"`
	}
	p := Struct{Field: "value"}

	insensitive, _, err := Detach(&p)
	assertNoError(t, err)
	assertTrue(t, insensitive != nil, "insensitive should be not nil")
	assertTrue(t, insensitive.Field == "", "sensitive field should be empty")
}

func TestDetachStructNotModified(t *testing.T) {
	t.Parallel()

	type Struct struct {
		Field string `sensitive:"true"`
	}
	p := Struct{Field: "value"}

	_, _, err := Detach(p)
	assertNoError(t, err)
	assertTrue(t, p.Field == "value", "original value should not be modified")
}

func TestDetachPointerModified(t *testing.T) {
	t.Parallel()

	type Struct struct {
		Field string `sensitive:"true"`
	}
	p := &Struct{Field: "value"}

	_, _, err := Detach(p)
	assertNoError(t, err)
	assertTrue(t, p.Field == "", "original value should be modified")
}
