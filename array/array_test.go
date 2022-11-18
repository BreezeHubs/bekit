package array

// func TestArrayForInt(t *testing.T) {
// 	arr := []int{1, 4, 5, 67, 89, 0, 346, 3}

// 	new, err := For(arr, func(val int, index int) any {
// 		if val > 10 {
// 			return val
// 		}
// 		return false
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(new)
// }

// func TestArrayForString(t *testing.T) {
// 	arr := []string{"1", "4", "5", "67", "89", "0", "346", "3"}

// 	new, err := For(arr, func(val string, index int) any {
// 		if len(val) > 2 {
// 			return val
// 		}
// 		return false
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(new)
// }

//func TestArrayFor(t *testing.T) {
//	testCases := []struct {
//		name    string
//		input   any
//		args    any
//		wantOut any
//		wantErr error
//	}{
//		{
//			name:  "array for int",
//			input: []int{1, 4, 5, 67, 89, 0, 346, 3},
//			args: func(val int, index int) any {
//				if val > 10 {
//					return val
//				}
//				return false
//			},
//			wantOut: []int{67, 89, 346},
//		},
//		{
//			name:  "array for string",
//			input: []string{"1", "4", "5", "67", "89", "0", "346", "3"},
//			args: func(val string, index int) any {
//				if len(val) > 2 {
//					return val
//				}
//				return false
//			},
//			wantOut: []string{"346"},
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			i := tc.input.([]any)
//			res, err := aa(i, tc.args)
//
//			assert.Equal(t, tc.wantErr, err)
//			if err != nil {
//				return
//			}
//
//			assert.Equal(t, tc.wantOut, res)
//		})
//	}
//}
//
//func aa[T arrayT](input []any, args any) ([]T, error) {
//	var i []T
//	for _, a := range input {
//		i = append(i, a)
//	}
//	a := args.(func(T, int) any)
//	return For(i, a)
//}
