package service

import (
	"context"
	"os"
	"testing"

	"g.hz.netease.com/horizon/core/middleware/user"
	"g.hz.netease.com/horizon/lib/orm"
	groupmanagermock "g.hz.netease.com/horizon/mock/pkg/group/manager"
	rolemock "g.hz.netease.com/horizon/mock/pkg/rbac/role"
	userauth "g.hz.netease.com/horizon/pkg/authentication/user"
	groupModels "g.hz.netease.com/horizon/pkg/group/models"
	"g.hz.netease.com/horizon/pkg/member"
	"g.hz.netease.com/horizon/pkg/member/models"
	roleservice "g.hz.netease.com/horizon/pkg/rbac/role"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	ctx context.Context
	s   Service
)

func PostMemberEqualsMember(postMember PostMember, member *models.Member) bool {
	return models.ResourceType(postMember.ResourceType) == member.ResourceType &&
		postMember.ResourceID == member.ResourceID &&
		postMember.MemberInfo == member.MemberNameID &&
		postMember.MemberType == member.MemberType &&
		postMember.Role == member.Role
}

// nolint
func TestCreateAndUpdateGroupMember(t *testing.T) {
	// mock the groupManager
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	groupManager := groupmanagermock.NewMockManager(mockCtrl)
	roleMockService := rolemock.NewMockService(mockCtrl)
	originService := &service{
		memberManager: member.Mgr,
		groupManager:  groupManager,
		roleService:   roleMockService,
	}
	s = originService

	//  case  /group1/group2
	//    group1 member: tom(1), jerry(1), cat(1)
	//    group2 member: tom(2), jerry(2)
	var group1ID uint = 3
	var group2ID uint = 4
	var traversalIDs string = "3,4"
	var tomID uint = 1
	var jerryID uint = 2
	var catID uint = 3
	var grandUser userauth.User = &userauth.DefaultInfo{
		Name:     "tom",
		FullName: "tom",
		ID:       tomID,
	}
	var ctx = context.WithValue(ctx, user.Key(), grandUser)
	// insert service to group2
	postMemberTom2 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group2ID,
		MemberInfo:   tomID,
		MemberType:   models.MemberUser,
		Role:         "owner",
	}
	member, err := originService.createMemberDirect(ctx, postMemberTom2)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberTom2, member))

	postMemberJerry2 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group2ID,
		MemberInfo:   jerryID,
		MemberType:   models.MemberUser,
		Role:         "owner",
	}
	member, err = originService.createMemberDirect(ctx, postMemberJerry2)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberJerry2, member))

	// insert member to group1
	postMemberTom1 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group1ID,
		MemberInfo:   tomID,
		MemberType:   models.MemberUser,
		Role:         "owner",
	}
	tomMember1, err := originService.createMemberDirect(ctx, postMemberTom1)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberTom1, tomMember1))

	postMemberJerry1 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group1ID,
		MemberInfo:   jerryID,
		MemberType:   models.MemberUser,
		Role:         "maintainer",
	}
	member, err = originService.createMemberDirect(ctx, postMemberJerry1)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberJerry1, member))

	postMemberCat1 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group1ID,
		MemberInfo:   catID,
		MemberType:   models.MemberUser,
		Role:         "maintainer",
	}
	member, err = originService.createMemberDirect(ctx, postMemberCat1)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberCat1, member))

	// test role smaller
	// create member exist  auto change to update role
	roleMockService.EXPECT().RoleCompare(ctx, gomock.Any(), gomock.Any()).Return(
		roleservice.RoleSmaller, nil).Times(1)
	// create member ok
	groupManager.EXPECT().GetByID(gomock.Any(),
		gomock.Any()).DoAndReturn(func(_ context.Context, id uint) (*groupModels.Group, error) {
		return &groupModels.Group{
			Model:           gorm.Model{},
			Name:            "",
			Path:            "",
			VisibilityLevel: "",
			Description:     "",
			ParentID:        0,
			TraversalIDs:    traversalIDs,
		}, nil
	}).Times(1)
	postMemberCat2 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group2ID,
		MemberInfo:   catID,
		MemberType:   models.MemberUser,
		Role:         "maintainer",
	}
	_, err = s.CreateMember(ctx, postMemberCat2)
	assert.Equal(t, err, ErrGrantHighRole)

	// create member exist  auto change to update role
	roleMockService.EXPECT().RoleCompare(ctx, gomock.Any(), gomock.Any()).Return(
		roleservice.RoleBigger, nil).AnyTimes()
	// create member ok
	groupManager.EXPECT().GetByID(gomock.Any(),
		gomock.Any()).DoAndReturn(func(_ context.Context, id uint) (*groupModels.Group, error) {
		return &groupModels.Group{
			Model:           gorm.Model{},
			Name:            "",
			Path:            "",
			VisibilityLevel: "",
			Description:     "",
			ParentID:        0,
			TraversalIDs:    traversalIDs,
		}, nil
	}).Times(1)
	postMemberCat2 = PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group2ID,
		MemberInfo:   catID,
		MemberType:   models.MemberUser,
		Role:         "maintainer",
	}
	catMember2, err := s.CreateMember(ctx, postMemberCat2)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberCat2, catMember2))

	groupManager.EXPECT().GetByID(gomock.Any(),
		gomock.Any()).DoAndReturn(func(_ context.Context, id uint) (*groupModels.Group, error) {
		return &groupModels.Group{
			Model:           gorm.Model{},
			Name:            "",
			Path:            "",
			VisibilityLevel: "",
			Description:     "",
			ParentID:        0,
			TraversalIDs:    traversalIDs,
		}, nil
	}).Times(1)
	postMemberCat2.Role = "develop"
	member, err = s.CreateMember(ctx, postMemberCat2)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberCat2, member))

	// update member not exist
	var memberIDNotExist uint = 123233434
	member, err = s.UpdateMember(ctx, memberIDNotExist, "owner")
	assert.Equal(t, err.Error(), ErrMemberNotExist.Error())

	// update member correct
	groupManager.EXPECT().GetByID(gomock.Any(),
		gomock.Any()).DoAndReturn(func(_ context.Context, id uint) (*groupModels.Group, error) {
		return &groupModels.Group{
			Model:           gorm.Model{},
			Name:            "",
			Path:            "",
			VisibilityLevel: "",
			Description:     "",
			ParentID:        0,
			TraversalIDs:    traversalIDs,
		}, nil
	}).Times(1)
	member, err = s.UpdateMember(ctx, tomMember1.ID, "maintainer")
	assert.Nil(t, err)
	assert.Equal(t, member.Role, "maintainer")
	assert.Equal(t, member.ID, tomMember1.ID)

	// remove member not exist
	err = s.RemoveMember(ctx, memberIDNotExist)
	assert.Equal(t, err.Error(), ErrMemberNotExist.Error())

	// remove member ok
	groupManager.EXPECT().GetByID(gomock.Any(),
		gomock.Any()).DoAndReturn(func(_ context.Context, id uint) (*groupModels.Group, error) {
		return &groupModels.Group{
			Model:           gorm.Model{},
			Name:            "",
			Path:            "",
			VisibilityLevel: "",
			Description:     "",
			ParentID:        0,
			TraversalIDs:    traversalIDs,
		}, nil
	}).Times(1)
	err = s.RemoveMember(ctx, catMember2.ID)
	assert.Nil(t, err)
}

// nolint
func TestListGroupMember(t *testing.T) {
	// mock the groupManager
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	groupManager := groupmanagermock.NewMockManager(mockCtrl)
	originService := &service{
		memberManager: member.Mgr,
		groupManager:  groupManager,
	}
	s = originService

	//  case  /group1/group2
	//    group1 service: tom(1), jerry(1), cat(1)
	//    group2 service: tom(2), jerry(2)
	//    ret: tom(2), jerry(2), cat(1)
	var group2ID uint = 2
	var group1ID uint = 1
	var traversalIDs string = "1,2"
	var tomID uint = 1
	var jerryID uint = 2
	var catID uint = 3
	var grandUser userauth.User = &userauth.DefaultInfo{
		Name:     "tom",
		FullName: "tom",
		ID:       tomID,
	}
	ctx = context.WithValue(ctx, user.Key(), grandUser)

	// insert service to group2
	postMemberTom2 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group2ID,
		MemberInfo:   tomID,
		MemberType:   models.MemberUser,
		Role:         "owner",
	}
	member, err := originService.createMemberDirect(ctx, postMemberTom2)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberTom2, member))

	postMemberJerry2 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group2ID,
		MemberInfo:   jerryID,
		MemberType:   models.MemberUser,
		Role:         "owner",
	}
	member, err = originService.createMemberDirect(ctx, postMemberJerry2)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberJerry2, member))

	// insert service to group1
	postMemberTom1 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group1ID,
		MemberInfo:   tomID,
		MemberType:   models.MemberUser,
		Role:         "owner",
	}
	member, err = originService.createMemberDirect(ctx, postMemberTom1)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberTom1, member))

	postMemberJerry1 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group1ID,
		MemberInfo:   jerryID,
		MemberType:   models.MemberUser,
		Role:         "maintainer",
	}
	member, err = originService.createMemberDirect(ctx, postMemberJerry1)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberJerry1, member))

	postMemberCat1 := PostMember{
		ResourceType: models.TypeGroupStr,
		ResourceID:   group1ID,
		MemberInfo:   catID,
		MemberType:   models.MemberUser,
		Role:         "maintainer",
	}
	member, err = originService.createMemberDirect(ctx, postMemberCat1)
	assert.Nil(t, err)
	assert.True(t, PostMemberEqualsMember(postMemberCat1, member))

	// listmember of group2
	groupManager.EXPECT().GetByID(gomock.Any(),
		gomock.Any()).DoAndReturn(func(_ context.Context, id uint) (*groupModels.Group, error) {
		return &groupModels.Group{
			Model:           gorm.Model{},
			Name:            "",
			Path:            "",
			VisibilityLevel: "",
			Description:     "",
			ParentID:        0,
			TraversalIDs:    traversalIDs,
		}, nil
	}).Times(1)
	members, err := s.ListMember(ctx, models.TypeGroupStr, group2ID)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(members))
	assert.True(t, PostMemberEqualsMember(postMemberTom2, &members[0]))
	assert.True(t, PostMemberEqualsMember(postMemberJerry2, &members[1]))
	assert.True(t, PostMemberEqualsMember(postMemberCat1, &members[2]))
}

func TestListApplicationMember(t *testing.T) {
	// TODO(tom)
}

func TestListApplicationInstanceMember(t *testing.T) {
	// TODO(tom)
}

func TestMain(m *testing.M) {
	db, _ := orm.NewSqliteDB("")
	if err := db.AutoMigrate(&models.Member{}); err != nil {
		panic(err)
	}

	ctx = orm.NewContext(context.TODO(), db)
	os.Exit(m.Run())
}