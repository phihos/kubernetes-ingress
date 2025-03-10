package api

import "github.com/haproxytech/client-native/v6/models"

func (c *clientNative) LogTargetCreate(id int64, parentType, parentName string, rule models.LogTarget) error {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return err
	}
	return configuration.CreateLogTarget(id, parentType, parentName, &rule, c.activeTransaction, 0)
}

func (c *clientNative) LogTargetDeleteAll(parentType, parentName string) (err error) {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return
	}
	_, rules, err := configuration.GetLogTargets(parentType, parentName, c.activeTransaction)
	if err != nil {
		return
	}
	for range rules {
		if err = configuration.DeleteLogTarget(0, parentType, parentName, c.activeTransaction, 0); err != nil {
			break
		}
	}
	return
}

func (c *clientNative) LogTargetsGet(parentType, parentName string) (models.LogTargets, error) {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return nil, err
	}

	_, rules, err := configuration.GetLogTargets(parentType, parentName, c.activeTransaction)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (c *clientNative) LogTargetsReplace(parentType, parentName string, rules models.LogTargets) error {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return err
	}

	err = configuration.ReplaceLogTargets(parentType, parentName, rules, c.activeTransaction, 0)
	if err != nil {
		return err
	}
	return nil
}
