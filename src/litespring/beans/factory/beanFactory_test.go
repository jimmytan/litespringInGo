package factory

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestDefaultBeanFactory(t *testing.T) {
	var factory = DefaultBeanFactory{beanDefinition: DefMap{}}

	var beanDefinitions = []DefaultBeanDefinition{
		{
			Name: "1",
			Build: func() (interface{}, error) {return nil, nil},
		},
		{
			Name: "2",
			Build: func() (interface{}, error) {return nil, nil},	
		},
	}
	
	s := make([]BeanDefinition, len(beanDefinitions))
	for i, v := range beanDefinitions {
		s[i] = v
	}

	factory.AddBeanDefinitions(s)
	require.Len(t, factory.beanDefinition, 2 )

}