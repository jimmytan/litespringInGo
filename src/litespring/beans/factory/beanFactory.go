package factory

import (
	"fmt"
	"errors"
)

type BeanFactory interface {
	GetBeanDefinition(beanID string) BeanDefinition
	GetBean(beanID string) 
	AddBeanDefinitions(definitions []BeanDefinition)
}

type DefaultBeanFactory struct {
	beanDefinition DefMap
}


func (factory *DefaultBeanFactory) AddBeanDefinitions(definitions []BeanDefinition) error {
	if(len(definitions) == 0){
		return errors.New("bean definitions is empty")
	}

	for _, def := range definitions {
		if _, ok := factory.beanDefinition[def.GetName()]; ok {
			fmt.Errorf("definition with name %s already exists", def.GetName())
		}else {
			factory.beanDefinition[def.GetName()] = def
		}
		
	}
	return nil
}

func (factory *DefaultBeanFactory) GetBeanDefinition(beanID string) BeanDefinition {
	return factory.beanDefinition[beanID]
}



