package modules

// item -> owner: userId, name, desc and role: role1, role2, role3, subitems: subitem1, subitem2, subitem3
// subitem -> settings {}, type

// role -> owner, permissions: {key: value}, prefix (root, customs)
// (на фронте будет кнопка "дать такие же права человеку" и будут обычные default, manager, owner)
// middleware (token-checker - написать и через grpc принимать - в токене требовать овнерство проектом) - addRoleToUser, removeUserRole, updateUserRole

// userinfo (user) - in auth

// image store

// IN NEXT VERSIONS - stats ( metrics )

// docker images : psql, auth, role, item
