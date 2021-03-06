// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "example.com/banana/teamcity/model"

// ITeamcityClient is an autogenerated mock type for the ITeamcityClient type
type ITeamcityClient struct {
	mock.Mock
}

// FetchBuildTypes provides a mock function with given fields:
func (_m *ITeamcityClient) FetchBuildTypes() *model.BuildTypes {
	ret := _m.Called()

	var r0 *model.BuildTypes
	if rf, ok := ret.Get(0).(func() *model.BuildTypes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BuildTypes)
		}
	}

	return r0
}

// FetchProjects provides a mock function with given fields:
func (_m *ITeamcityClient) FetchProjects() *model.Projects {
	ret := _m.Called()

	var r0 *model.Projects
	if rf, ok := ret.Get(0).(func() *model.Projects); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Projects)
		}
	}

	return r0
}
