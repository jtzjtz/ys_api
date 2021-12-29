package entity

type UserBlack struct {
	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BlackUserId int32  `protobuf:"varint,3,opt,name=black_user_id,json=blackUserId,proto3" json:"black_user_id,omitempty"`
	Status      int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}
