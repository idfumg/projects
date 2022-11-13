package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	scenarios := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with one non-string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Moscow"},
			},
			[]string{"London", "Moscow"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Moscow"},
			},
			[]string{"London", "Moscow"},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			var got []string

			walk(scenario.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, scenario.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, scenario.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) { // because of the random order of hash elements
		m := map[string]string{
			"Foo": "123",
			"Bar": "321",
		}
		var got []string
		walk(m, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "123")
		assertContains(t, got, "321")
	})

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan Profile)

		go func(){
			ch <- Profile{33, "Berlin"}
			ch <- Profile{34, "Moscow"}
			close(ch)
		}()

		var got []string
		want := []string{"Berlin", "Moscow"}

		walk(ch, func(input string){
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("with function", func(t *testing.T) {
		fn := func()(Profile, Profile){
			return Profile{33, "Berlin"}, Profile{34, "Moscow"}
		}

		var got []string
		want := []string{"Berlin", "Moscow"}

		walk(fn, func(input string){
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, arr []string, value string) {
	t.Helper()
	for _, x := range arr {
		if x == value {
			return
		}
	}
	t.Errorf("expected %+v to contain %q but it didn't", arr, value)
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		results := val.Call(nil)
		for _, result := range results {
			walk(result.Interface(), fn)
		}
	}
}
