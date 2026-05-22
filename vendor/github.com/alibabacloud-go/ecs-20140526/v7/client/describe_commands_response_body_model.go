// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommands(v *DescribeCommandsResponseBodyCommands) *DescribeCommandsResponseBody
	GetCommands() *DescribeCommandsResponseBodyCommands
	SetNextToken(v string) *DescribeCommandsResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *DescribeCommandsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCommandsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeCommandsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeCommandsResponseBody
	GetTotalCount() *int64
}

type DescribeCommandsResponseBody struct {
	Commands *DescribeCommandsResponseBodyCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of commands.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCommandsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBody) GetCommands() *DescribeCommandsResponseBodyCommands {
	return s.Commands
}

func (s *DescribeCommandsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCommandsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCommandsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCommandsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommandsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCommandsResponseBody) SetCommands(v *DescribeCommandsResponseBodyCommands) *DescribeCommandsResponseBody {
	s.Commands = v
	return s
}

func (s *DescribeCommandsResponseBody) SetNextToken(v string) *DescribeCommandsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetPageNumber(v int64) *DescribeCommandsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetPageSize(v int64) *DescribeCommandsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetRequestId(v string) *DescribeCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetTotalCount(v int64) *DescribeCommandsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCommandsResponseBody) Validate() error {
	if s.Commands != nil {
		if err := s.Commands.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCommandsResponseBodyCommands struct {
	Command []*DescribeCommandsResponseBodyCommandsCommand `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s DescribeCommandsResponseBodyCommands) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommands) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommands) GetCommand() []*DescribeCommandsResponseBodyCommandsCommand {
	return s.Command
}

func (s *DescribeCommandsResponseBodyCommands) SetCommand(v []*DescribeCommandsResponseBodyCommandsCommand) *DescribeCommandsResponseBodyCommands {
	s.Command = v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) Validate() error {
	if s.Command != nil {
		for _, item := range s.Command {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommandsResponseBodyCommandsCommand struct {
	Category             *string                                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	CommandContent       *string                                                          `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandId            *string                                                          `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	CreationTime         *string                                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description          *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableParameter      *bool                                                            `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	InvokeTimes          *int32                                                           `json:"InvokeTimes,omitempty" xml:"InvokeTimes,omitempty"`
	Latest               *bool                                                            `json:"Latest,omitempty" xml:"Latest,omitempty"`
	Launcher             *string                                                          `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	Name                 *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	ParameterDefinitions *DescribeCommandsResponseBodyCommandsCommandParameterDefinitions `json:"ParameterDefinitions,omitempty" xml:"ParameterDefinitions,omitempty" type:"Struct"`
	ParameterNames       *DescribeCommandsResponseBodyCommandsCommandParameterNames       `json:"ParameterNames,omitempty" xml:"ParameterNames,omitempty" type:"Struct"`
	Provider             *string                                                          `json:"Provider,omitempty" xml:"Provider,omitempty"`
	ResourceGroupId      *string                                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags                 *DescribeCommandsResponseBodyCommandsCommandTags                 `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Timeout              *int64                                                           `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type                 *string                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	Version              *int32                                                           `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkingDir           *string                                                          `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeCommandsResponseBodyCommandsCommand) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsCommand) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetCategory() *string {
	return s.Category
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetCommandContent() *string {
	return s.CommandContent
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetDescription() *string {
	return s.Description
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetInvokeTimes() *int32 {
	return s.InvokeTimes
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetLatest() *bool {
	return s.Latest
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetLauncher() *string {
	return s.Launcher
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetName() *string {
	return s.Name
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetParameterDefinitions() *DescribeCommandsResponseBodyCommandsCommandParameterDefinitions {
	return s.ParameterDefinitions
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetParameterNames() *DescribeCommandsResponseBodyCommandsCommandParameterNames {
	return s.ParameterNames
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetProvider() *string {
	return s.Provider
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetTags() *DescribeCommandsResponseBodyCommandsCommandTags {
	return s.Tags
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetTimeout() *int64 {
	return s.Timeout
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetType() *string {
	return s.Type
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeCommandsResponseBodyCommandsCommand) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetCategory(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.Category = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetCommandContent(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.CommandContent = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetCommandId(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetCreationTime(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.CreationTime = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetDescription(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.Description = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetEnableParameter(v bool) *DescribeCommandsResponseBodyCommandsCommand {
	s.EnableParameter = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetInvokeTimes(v int32) *DescribeCommandsResponseBodyCommandsCommand {
	s.InvokeTimes = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetLatest(v bool) *DescribeCommandsResponseBodyCommandsCommand {
	s.Latest = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetLauncher(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.Launcher = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetName(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.Name = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetParameterDefinitions(v *DescribeCommandsResponseBodyCommandsCommandParameterDefinitions) *DescribeCommandsResponseBodyCommandsCommand {
	s.ParameterDefinitions = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetParameterNames(v *DescribeCommandsResponseBodyCommandsCommandParameterNames) *DescribeCommandsResponseBodyCommandsCommand {
	s.ParameterNames = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetProvider(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.Provider = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetResourceGroupId(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetTags(v *DescribeCommandsResponseBodyCommandsCommandTags) *DescribeCommandsResponseBodyCommandsCommand {
	s.Tags = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetTimeout(v int64) *DescribeCommandsResponseBodyCommandsCommand {
	s.Timeout = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetType(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.Type = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetVersion(v int32) *DescribeCommandsResponseBodyCommandsCommand {
	s.Version = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) SetWorkingDir(v string) *DescribeCommandsResponseBodyCommandsCommand {
	s.WorkingDir = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommand) Validate() error {
	if s.ParameterDefinitions != nil {
		if err := s.ParameterDefinitions.Validate(); err != nil {
			return err
		}
	}
	if s.ParameterNames != nil {
		if err := s.ParameterNames.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCommandsResponseBodyCommandsCommandParameterDefinitions struct {
	ParameterDefinition []*DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition `json:"ParameterDefinition,omitempty" xml:"ParameterDefinition,omitempty" type:"Repeated"`
}

func (s DescribeCommandsResponseBodyCommandsCommandParameterDefinitions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsCommandParameterDefinitions) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitions) GetParameterDefinition() []*DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition {
	return s.ParameterDefinition
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitions) SetParameterDefinition(v []*DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) *DescribeCommandsResponseBodyCommandsCommandParameterDefinitions {
	s.ParameterDefinition = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitions) Validate() error {
	if s.ParameterDefinition != nil {
		for _, item := range s.ParameterDefinition {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition struct {
	DefaultValue   *string                                                                                           `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description    *string                                                                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	ParameterName  *string                                                                                           `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	PatternRegex   *string                                                                                           `json:"PatternRegex,omitempty" xml:"PatternRegex,omitempty"`
	PossibleValues *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues `json:"PossibleValues,omitempty" xml:"PossibleValues,omitempty" type:"Struct"`
	Required       *bool                                                                                             `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) GetDescription() *string {
	return s.Description
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) GetPatternRegex() *string {
	return s.PatternRegex
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) GetPossibleValues() *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues {
	return s.PossibleValues
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) GetRequired() *bool {
	return s.Required
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) SetDefaultValue(v string) *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition {
	s.DefaultValue = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) SetDescription(v string) *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition {
	s.Description = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) SetParameterName(v string) *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition {
	s.ParameterName = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) SetPatternRegex(v string) *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition {
	s.PatternRegex = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) SetPossibleValues(v *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues) *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition {
	s.PossibleValues = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) SetRequired(v bool) *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition {
	s.Required = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinition) Validate() error {
	if s.PossibleValues != nil {
		if err := s.PossibleValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues struct {
	PossibleValue []*string `json:"PossibleValue,omitempty" xml:"PossibleValue,omitempty" type:"Repeated"`
}

func (s DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues) GetPossibleValue() []*string {
	return s.PossibleValue
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues) SetPossibleValue(v []*string) *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues {
	s.PossibleValue = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterDefinitionsParameterDefinitionPossibleValues) Validate() error {
	return dara.Validate(s)
}

type DescribeCommandsResponseBodyCommandsCommandParameterNames struct {
	ParameterName []*string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty" type:"Repeated"`
}

func (s DescribeCommandsResponseBodyCommandsCommandParameterNames) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsCommandParameterNames) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterNames) GetParameterName() []*string {
	return s.ParameterName
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterNames) SetParameterName(v []*string) *DescribeCommandsResponseBodyCommandsCommandParameterNames {
	s.ParameterName = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandParameterNames) Validate() error {
	return dara.Validate(s)
}

type DescribeCommandsResponseBodyCommandsCommandTags struct {
	Tag []*DescribeCommandsResponseBodyCommandsCommandTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCommandsResponseBodyCommandsCommandTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsCommandTags) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsCommandTags) GetTag() []*DescribeCommandsResponseBodyCommandsCommandTagsTag {
	return s.Tag
}

func (s *DescribeCommandsResponseBodyCommandsCommandTags) SetTag(v []*DescribeCommandsResponseBodyCommandsCommandTagsTag) *DescribeCommandsResponseBodyCommandsCommandTags {
	s.Tag = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommandsResponseBodyCommandsCommandTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeCommandsResponseBodyCommandsCommandTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsCommandTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsCommandTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeCommandsResponseBodyCommandsCommandTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeCommandsResponseBodyCommandsCommandTagsTag) SetTagKey(v string) *DescribeCommandsResponseBodyCommandsCommandTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandTagsTag) SetTagValue(v string) *DescribeCommandsResponseBodyCommandsCommandTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsCommandTagsTag) Validate() error {
	return dara.Validate(s)
}
