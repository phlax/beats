// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]bool

func init() {
	// Packed Files
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/metricbeat.yml
	unpacked := packer.MustUnpack("eJzsWE2To7gZvudnzDWpBMTghFTtwdCLALvpMe5GQjckuQFbgKvBH5DKf08JDDb27O7MZDenHLragD7ej+d93kf616dqv2F/2xR8X2ZF/dcmF5/++Ynmdk1eyyRA+o5BY0+L1YIVYUXw88zKlITk4jPWfIXlYUpfy2SjKYvL+4rDsMEaSakTCtaWC9eaJ67jCwrDLYdG85KZe1qYKnee+2+wFpvXMuGOOJFVmch3PLcrjsL2Jemfl2tT0Dw4USAOfF4ulmszj9FZIdh7vx/LQFgR5CtUc2cuFAfmhAo79d9iEOovmVzLzigMd1Y2T1zLBBE6q71v8nleu07QcPQ28elxfe/WvpzCUPD5+Dz10VISqoWKtMVK9loMw8NLZroR9k1a7EoPkIqgUPGaXeJlUbKGdrvGgaD4ufRwTWOknzgOWoyVvWtFuftzmjJFTykKWwbtLVknxbC+15wSD4RVhH0lRn5LkN1EICkWq/KnT3/p0/2eiQ3dxA/pzkVFkCcivFpEQN29ZF2YRZSHKZ/vU5bz9iUzqZuptpudErfwBXfC0zIXFV3rY1i/IJl+X3Rj7scW0i2zinAglnl4iJBXEbQySG5XDLxlS2ueLd/6/xTZhwhxQVF44JZeUxCILzippctxo/Yht9zKtdw6WMv/Xh0hPSUgrAnSldv1ueOpZD0ZW1HAixjpxTI/C56H1RcUiKgIC1covwn3LsX5XkRa8B4jfUdwcoGPqWywKV4yU67fWklZSxjGebjlttFyxxMRUt+Z4x0jELYMGCO0KNDfI2AcSH7eR9qqhy8IG24bKSkCwS7julJ5LRMKjYKdunKoIzwf9j9FOCgfyszxVAonJbWlmqljYFfUNhSqGlWMfWX43s8Jji+ZOdjcXsrjF2zVj1SbL8a5lik20BfMWc1cu8qYFjQE2TVrutjvSZc3b8RMjxGzocAXTPOPrHjO8OrOVi04YnDeM201u7UllqXyWia3MR1K/vv9GGOeEUSOLO9pgHVYUsUQH1J4R7nnUJpD7ml7fTf63OdIbJygiZCv9D4E73Ff4r+QN//OXnPPYVg/+jHds8ebmrKngXrNlMPkgQqv+L5SXoT0lkJbIa/lbR4HuzpcT2KH7BO7ocbxvcRl2+PzdrxrdTXRcUpfH2xxO8+FIejqFpA9hW8zF56P5DRdm2AiaLE6Mi3YxejzzHV8hUBxuNtH1vOBW4YSafOZ6wRbNv/6OhwFp5fMVIkzv7PlfCSNsaPA/yBYtpLgGIF6rMFbX5eaJwgULdb8impc+jVznf7do//syDTRynkylxvs38Shw1lNNSIw6NribR4+CN7NXCdsSWgcOQ5OHK++fR4MUg7tlvzYnEPXI7B/ipAvrpjxBEUGIKHRjZ/UZN9qdxEOUgZCZZnLdvX2R8uEaSx7jIzPMtc3eB8xsVybQ55vY7OL8fPNWn3OrvOf2/E31AWHpKLade/YCRQG9y0F1/1p+wx8ZKsECuU6155x7Ikej0HJ0dWGCJwlPnOqhQfiePp1Dj9SwCuCTTUqfDW6+v8RIfIRrf63EikCxmnEDyQNBcrDuhz+Y+ZCWddhRpCt9DxF9tQJBNs+5n/EhNXjZdAlD3JqjEktNriTZcIqeEkkN3QySSgEqe3CCv+OwHX/ZX6xT1WNxXqee70ce5EczzQzjcBb6QFD5Y6pcotVrsUFhXbLodgyEKYs98tOal1q32tO0rYiBnYeg5+LpTUvOh7UJP6n8ivf1B8Z+4oAe0WhwnKxvQiuLUWyIamCO94+Ahdhhr2OXFEzCpmW4EBllr6nUDkQpKdRfhbkSc0jdG7vRM8wVgbl1JH9bwm3XE1pbhcEqZIwDxQZO/Kqfl5iGaSqJnIvvPo14XZdH3egmY6FRkGAOJBGr7oieFJ3BHkqaTxuFeaR9QKr5dCuaN/UZIM78jx874BxQ1wM2kr8VEqbj8TpGtyBWIYkWYUg5bBBauWOBG4qMpYEr2ZSOFIQdAJpma86gpbNdlmImlr6LsZ+H1PL/T2EYU1w0Eht/oeKwwuWLuKjkbig6G2wJWe5UT+KxE4oPQqQ3xIVQwzw/8Xm7yE2CU4VltsSE12cMOhJbCTwC86/IjSnNt6970l4JOavNgms8T2H6TvLw4Lg9DTseWkMNyLQbCTOcPb5YwkuHKU9Dw2jJfLgqAVHtr0XpxfSF8ZdU/mFOdmtIBkPdBec3QkI2eTlQW/1w4ciGfcDA+eUw/CdA6HEttEQxMXGmU9wMOTgVpBc1pQ1vlk8GasvTqAz+PbnZVbtH2PU/3V7PJWJ14+9FSCTA8ekuU7GXg8Bkps5Ol8FGlTTGITvEfaa6F6gXTAy8gQYm+kNVgab5eH9+TuE4828q8hqCOp5u8PSxK9e3I2963vmQFuJ8HfuM8yx9C7GPy5oL/cvq2/fm2Nf4EFY9HdZHd98y8GSTvTBr/L60GOmYmfCA2OP/J57rsm6V5yN/oxClGDZQ41G1iIHqaDb8uFea+JPf8HT1eN/c8+17GzgDUHBnjWdcOtF8br73V3KyENNbLG9lfz006d//+k/AQAA//97aRE7")
	SupportedMap = make(map[string]bool)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Name)] = true
	}
}
