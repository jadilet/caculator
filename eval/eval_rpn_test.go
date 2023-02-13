package eval

import (
	"container/list"
	"testing"
)

func TestEvalRpn(t *testing.T) {
	lst1 := list.New()
	lst1.PushBack("2")
	lst1.PushBack("1")
	lst1.PushBack("+")
	lst1.PushBack("3")
	lst1.PushBack("*")


	lst2 := list.New()
	lst2.PushBack("2")
	lst2.PushBack("1")
	lst2.PushBack("+")
	lst2.PushBack("*")

	lst3 := list.New()
	lst3.PushBack("2")
	lst3.PushBack("-")

	tests := []struct {
		name    string
		tokens  *list.List
		want    int
		wantErr bool
	}{
		{
			name:   "Valid Reverse Polish Notation",
			tokens: lst1,
			want:   9,
		},
		{
			name:    "Invalid Reverse Polish Notation",
			tokens:  lst2,
			wantErr: true,
		},
		{
			name:    "Invalid input reverse polish notation",
			tokens:  list.New(),
			wantErr: true,
		},
		{
			name:    "Error invalid input either negative literal or unary operation",
			tokens:  lst3,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalRpn(tt.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalRpn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EvalRpn() = %v, want %v", got, tt.want)
			}
		})
	}
}
