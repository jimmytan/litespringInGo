package factory

type DefaultBeanDefinition struct {
	Name string
	Build func() (interface{}, error)
}

type BeanDefinition interface {
	GetName() string
}

type DefMap map[string]BeanDefinition

func (d DefaultBeanDefinition) GetName() string {
	return d.Name
}