// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: pkg/common/proto/agent.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bet_Type int32

const (
	Bet_OVER  Bet_Type = 0
	Bet_UNDER Bet_Type = 1
)

// Enum value maps for Bet_Type.
var (
	Bet_Type_name = map[int32]string{
		0: "OVER",
		1: "UNDER",
	}
	Bet_Type_value = map[string]int32{
		"OVER":  0,
		"UNDER": 1,
	}
)

func (x Bet_Type) Enum() *Bet_Type {
	p := new(Bet_Type)
	*p = x
	return p
}

func (x Bet_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Bet_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_common_proto_agent_proto_enumTypes[0].Descriptor()
}

func (Bet_Type) Type() protoreflect.EnumType {
	return &file_pkg_common_proto_agent_proto_enumTypes[0]
}

func (x Bet_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Bet_Type.Descriptor instead.
func (Bet_Type) EnumDescriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{1, 0}
}

type Bot_Source int32

const (
	Bot_LOCAL  Bot_Source = 0
	Bot_REMOTE Bot_Source = 1
)

// Enum value maps for Bot_Source.
var (
	Bot_Source_name = map[int32]string{
		0: "LOCAL",
		1: "REMOTE",
	}
	Bot_Source_value = map[string]int32{
		"LOCAL":  0,
		"REMOTE": 1,
	}
)

func (x Bot_Source) Enum() *Bot_Source {
	p := new(Bot_Source)
	*p = x
	return p
}

func (x Bot_Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Bot_Source) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_common_proto_agent_proto_enumTypes[1].Descriptor()
}

func (Bot_Source) Type() protoreflect.EnumType {
	return &file_pkg_common_proto_agent_proto_enumTypes[1]
}

func (x Bot_Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Bot_Source.Descriptor instead.
func (Bot_Source) EnumDescriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{7, 0}
}

type FantasyLandscape struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchNumber uint32          `protobuf:"varint,1,opt,name=match_number,json=matchNumber,proto3" json:"match_number,omitempty"`
	Settings    *LeagueSettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	BotTeam     *FantasyTeam    `protobuf:"bytes,3,opt,name=bot_team,json=botTeam,proto3" json:"bot_team,omitempty"`
	Bet         *Bet            `protobuf:"bytes,4,opt,name=bet,proto3" json:"bet,omitempty"`
	Players     []*Player       `protobuf:"bytes,5,rep,name=players,proto3" json:"players,omitempty"` // TODO historical / statistical data
}

func (x *FantasyLandscape) Reset() {
	*x = FantasyLandscape{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FantasyLandscape) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FantasyLandscape) ProtoMessage() {}

func (x *FantasyLandscape) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FantasyLandscape.ProtoReflect.Descriptor instead.
func (*FantasyLandscape) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{0}
}

func (x *FantasyLandscape) GetMatchNumber() uint32 {
	if x != nil {
		return x.MatchNumber
	}
	return 0
}

func (x *FantasyLandscape) GetSettings() *LeagueSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *FantasyLandscape) GetBotTeam() *FantasyTeam {
	if x != nil {
		return x.BotTeam
	}
	return nil
}

func (x *FantasyLandscape) GetBet() *Bet {
	if x != nil {
		return x.Bet
	}
	return nil
}

func (x *FantasyLandscape) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type Bet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfessionalHomeTeam string   `protobuf:"bytes,1,opt,name=professional_home_team,json=professionalHomeTeam,proto3" json:"professional_home_team,omitempty"`
	ProfessionalAwayTeam string   `protobuf:"bytes,2,opt,name=professional_away_team,json=professionalAwayTeam,proto3" json:"professional_away_team,omitempty"`
	Player               *Player  `protobuf:"bytes,3,opt,name=player,proto3" json:"player,omitempty"`
	Type                 Bet_Type `protobuf:"varint,4,opt,name=type,proto3,enum=Bet_Type" json:"type,omitempty"`
	Points               float32  `protobuf:"fixed32,5,opt,name=points,proto3" json:"points,omitempty"`
	Price                float32  `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Bet) Reset() {
	*x = Bet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bet) ProtoMessage() {}

func (x *Bet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bet.ProtoReflect.Descriptor instead.
func (*Bet) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{1}
}

func (x *Bet) GetProfessionalHomeTeam() string {
	if x != nil {
		return x.ProfessionalHomeTeam
	}
	return ""
}

func (x *Bet) GetProfessionalAwayTeam() string {
	if x != nil {
		return x.ProfessionalAwayTeam
	}
	return ""
}

func (x *Bet) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *Bet) GetType() Bet_Type {
	if x != nil {
		return x.Type
	}
	return Bet_OVER
}

func (x *Bet) GetPoints() float32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *Bet) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type FantasySelections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MakeBet bool          `protobuf:"varint,1,opt,name=make_bet,json=makeBet,proto3" json:"make_bet,omitempty"`
	Slots   []*PlayerSlot `protobuf:"bytes,2,rep,name=slots,proto3" json:"slots,omitempty"`
}

func (x *FantasySelections) Reset() {
	*x = FantasySelections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FantasySelections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FantasySelections) ProtoMessage() {}

func (x *FantasySelections) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FantasySelections.ProtoReflect.Descriptor instead.
func (*FantasySelections) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{2}
}

func (x *FantasySelections) GetMakeBet() bool {
	if x != nil {
		return x.MakeBet
	}
	return false
}

func (x *FantasySelections) GetSlots() []*PlayerSlot {
	if x != nil {
		return x.Slots
	}
	return nil
}

type FantasyTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *FantasyTeam) Reset() {
	*x = FantasyTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FantasyTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FantasyTeam) ProtoMessage() {}

func (x *FantasyTeam) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FantasyTeam.ProtoReflect.Descriptor instead.
func (*FantasyTeam) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{3}
}

func (x *FantasyTeam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FantasyTeam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FantasyTeam) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type LeagueSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumTeams     uint32        `protobuf:"varint,1,opt,name=num_teams,json=numTeams,proto3" json:"num_teams,omitempty"`
	SlotsPerTeam []*PlayerSlot `protobuf:"bytes,2,rep,name=slots_per_team,json=slotsPerTeam,proto3" json:"slots_per_team,omitempty"` // scoring, budgets, etc.
}

func (x *LeagueSettings) Reset() {
	*x = LeagueSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeagueSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeagueSettings) ProtoMessage() {}

func (x *LeagueSettings) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeagueSettings.ProtoReflect.Descriptor instead.
func (*LeagueSettings) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{4}
}

func (x *LeagueSettings) GetNumTeams() uint32 {
	if x != nil {
		return x.NumTeams
	}
	return 0
}

func (x *LeagueSettings) GetSlotsPerTeam() []*PlayerSlot {
	if x != nil {
		return x.SlotsPerTeam
	}
	return nil
}

type PlayerSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AllowedPlayerPositions []string `protobuf:"bytes,2,rep,name=allowed_player_positions,json=allowedPlayerPositions,proto3" json:"allowed_player_positions,omitempty"`
	AssignedPlayerId       string   `protobuf:"bytes,3,opt,name=assigned_player_id,json=assignedPlayerId,proto3" json:"assigned_player_id,omitempty"`
}

func (x *PlayerSlot) Reset() {
	*x = PlayerSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerSlot) ProtoMessage() {}

func (x *PlayerSlot) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerSlot.ProtoReflect.Descriptor instead.
func (*PlayerSlot) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerSlot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerSlot) GetAllowedPlayerPositions() []string {
	if x != nil {
		return x.AllowedPlayerPositions
	}
	return nil
}

func (x *PlayerSlot) GetAssignedPlayerId() string {
	if x != nil {
		return x.AssignedPlayerId
	}
	return ""
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName         string   `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	AllowedPositions []string `protobuf:"bytes,3,rep,name=allowed_positions,json=allowedPositions,proto3" json:"allowed_positions,omitempty"`
	ProfessionalTeam string   `protobuf:"bytes,4,opt,name=professional_team,json=professionalTeam,proto3" json:"professional_team,omitempty"`
	Status           string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	FantasyTeamId    uint32   `protobuf:"varint,6,opt,name=fantasy_team_id,json=fantasyTeamId,proto3" json:"fantasy_team_id,omitempty"` // 0 should be unassigned
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{6}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Player) GetAllowedPositions() []string {
	if x != nil {
		return x.AllowedPositions
	}
	return nil
}

func (x *Player) GetProfessionalTeam() string {
	if x != nil {
		return x.ProfessionalTeam
	}
	return ""
}

func (x *Player) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Player) GetFantasyTeamId() uint32 {
	if x != nil {
		return x.FantasyTeamId
	}
	return 0
}

type Bot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SourceType         Bot_Source `protobuf:"varint,2,opt,name=source_type,json=sourceType,proto3,enum=Bot_Source" json:"source_type,omitempty"`
	SourceRepoUsername string     `protobuf:"bytes,3,opt,name=source_repo_username,json=sourceRepoUsername,proto3" json:"source_repo_username,omitempty"`
	SourceRepoName     string     `protobuf:"bytes,4,opt,name=source_repo_name,json=sourceRepoName,proto3" json:"source_repo_name,omitempty"`
	SourcePath         string     `protobuf:"bytes,5,opt,name=source_path,json=sourcePath,proto3" json:"source_path,omitempty"`
	FantasyTeamId      uint32     `protobuf:"varint,6,opt,name=fantasy_team_id,json=fantasyTeamId,proto3" json:"fantasy_team_id,omitempty"` // 0 should be unassigned
}

func (x *Bot) Reset() {
	*x = Bot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_agent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bot) ProtoMessage() {}

func (x *Bot) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_agent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bot.ProtoReflect.Descriptor instead.
func (*Bot) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{7}
}

func (x *Bot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bot) GetSourceType() Bot_Source {
	if x != nil {
		return x.SourceType
	}
	return Bot_LOCAL
}

func (x *Bot) GetSourceRepoUsername() string {
	if x != nil {
		return x.SourceRepoUsername
	}
	return ""
}

func (x *Bot) GetSourceRepoName() string {
	if x != nil {
		return x.SourceRepoName
	}
	return ""
}

func (x *Bot) GetSourcePath() string {
	if x != nil {
		return x.SourcePath
	}
	return ""
}

func (x *Bot) GetFantasyTeamId() uint32 {
	if x != nil {
		return x.FantasyTeamId
	}
	return 0
}

type Simulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Landscape     *FantasyLandscape `protobuf:"bytes,2,opt,name=landscape,proto3" json:"landscape,omitempty"`
	NumIterations uint32            `protobuf:"varint,3,opt,name=num_iterations,json=numIterations,proto3" json:"num_iterations,omitempty"`
}

func (x *Simulation) Reset() {
	*x = Simulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_agent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simulation) ProtoMessage() {}

func (x *Simulation) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_agent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simulation.ProtoReflect.Descriptor instead.
func (*Simulation) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_agent_proto_rawDescGZIP(), []int{8}
}

func (x *Simulation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Simulation) GetLandscape() *FantasyLandscape {
	if x != nil {
		return x.Landscape
	}
	return nil
}

func (x *Simulation) GetNumIterations() uint32 {
	if x != nil {
		return x.NumIterations
	}
	return 0
}

var File_pkg_common_proto_agent_proto protoreflect.FileDescriptor

var file_pkg_common_proto_agent_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6,
	0x01, 0x0a, 0x10, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x79, 0x4c, 0x61, 0x6e, 0x64, 0x73, 0x63,
	0x61, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x79, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x07, 0x62, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x03,
	0x62, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x42, 0x65, 0x74, 0x52,
	0x03, 0x62, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0xfc, 0x01, 0x0a, 0x03, 0x42, 0x65, 0x74, 0x12,
	0x34, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x48, 0x6f, 0x6d,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x77, 0x61, 0x79, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x41, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x1f, 0x0a, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x42, 0x65, 0x74,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x1b, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x55,
	0x4e, 0x44, 0x45, 0x52, 0x10, 0x01, 0x22, 0x51, 0x0a, 0x11, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73,
	0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x61, 0x6b, 0x65, 0x5f, 0x62, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d,
	0x61, 0x6b, 0x65, 0x42, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6c,
	0x6f, 0x74, 0x52, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x22, 0x47, 0x0a, 0x0b, 0x46, 0x61, 0x6e,
	0x74, 0x61, 0x73, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x22, 0x60, 0x0a, 0x0e, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x65, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x54, 0x65, 0x61, 0x6d,
	0x73, 0x12, 0x31, 0x0a, 0x0e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x74,
	0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x0c, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x50, 0x65, 0x72,
	0x54, 0x65, 0x61, 0x6d, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xcf, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x61,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x61, 0x6e,
	0x74, 0x61, 0x73, 0x79, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x66, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x22, 0x89, 0x02, 0x0a, 0x03, 0x42, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x42, 0x6f, 0x74, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x79, 0x5f,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66,
	0x61, 0x6e, 0x74, 0x61, 0x73, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x06,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x10, 0x01, 0x22, 0x74, 0x0a,
	0x0a, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x6c,
	0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x79, 0x4c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70,
	0x65, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x6e, 0x75, 0x6d, 0x5f, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0x50, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x15, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x61,
	0x6e, 0x74, 0x61, 0x73, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x11, 0x2e, 0x46,
	0x61, 0x6e, 0x74, 0x61, 0x73, 0x79, 0x4c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x1a,
	0x12, 0x2e, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_common_proto_agent_proto_rawDescOnce sync.Once
	file_pkg_common_proto_agent_proto_rawDescData = file_pkg_common_proto_agent_proto_rawDesc
)

func file_pkg_common_proto_agent_proto_rawDescGZIP() []byte {
	file_pkg_common_proto_agent_proto_rawDescOnce.Do(func() {
		file_pkg_common_proto_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_common_proto_agent_proto_rawDescData)
	})
	return file_pkg_common_proto_agent_proto_rawDescData
}

var file_pkg_common_proto_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_common_proto_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_common_proto_agent_proto_goTypes = []interface{}{
	(Bet_Type)(0),             // 0: Bet.Type
	(Bot_Source)(0),           // 1: Bot.Source
	(*FantasyLandscape)(nil),  // 2: FantasyLandscape
	(*Bet)(nil),               // 3: Bet
	(*FantasySelections)(nil), // 4: FantasySelections
	(*FantasyTeam)(nil),       // 5: FantasyTeam
	(*LeagueSettings)(nil),    // 6: LeagueSettings
	(*PlayerSlot)(nil),        // 7: PlayerSlot
	(*Player)(nil),            // 8: Player
	(*Bot)(nil),               // 9: Bot
	(*Simulation)(nil),        // 10: Simulation
}
var file_pkg_common_proto_agent_proto_depIdxs = []int32{
	6,  // 0: FantasyLandscape.settings:type_name -> LeagueSettings
	5,  // 1: FantasyLandscape.bot_team:type_name -> FantasyTeam
	3,  // 2: FantasyLandscape.bet:type_name -> Bet
	8,  // 3: FantasyLandscape.players:type_name -> Player
	8,  // 4: Bet.player:type_name -> Player
	0,  // 5: Bet.type:type_name -> Bet.Type
	7,  // 6: FantasySelections.slots:type_name -> PlayerSlot
	7,  // 7: LeagueSettings.slots_per_team:type_name -> PlayerSlot
	1,  // 8: Bot.source_type:type_name -> Bot.Source
	2,  // 9: Simulation.landscape:type_name -> FantasyLandscape
	2,  // 10: AgentService.PerformFantasyActions:input_type -> FantasyLandscape
	4,  // 11: AgentService.PerformFantasyActions:output_type -> FantasySelections
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_common_proto_agent_proto_init() }
func file_pkg_common_proto_agent_proto_init() {
	if File_pkg_common_proto_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_common_proto_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FantasyLandscape); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_common_proto_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_common_proto_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FantasySelections); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_common_proto_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FantasyTeam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_common_proto_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeagueSettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_common_proto_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerSlot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_common_proto_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_common_proto_agent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_common_proto_agent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simulation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_common_proto_agent_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_common_proto_agent_proto_goTypes,
		DependencyIndexes: file_pkg_common_proto_agent_proto_depIdxs,
		EnumInfos:         file_pkg_common_proto_agent_proto_enumTypes,
		MessageInfos:      file_pkg_common_proto_agent_proto_msgTypes,
	}.Build()
	File_pkg_common_proto_agent_proto = out.File
	file_pkg_common_proto_agent_proto_rawDesc = nil
	file_pkg_common_proto_agent_proto_goTypes = nil
	file_pkg_common_proto_agent_proto_depIdxs = nil
}
