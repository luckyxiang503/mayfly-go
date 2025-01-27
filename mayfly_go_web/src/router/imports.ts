import RouterParent from '@/views/layout/routerView/parent.vue';

export const imports = {
    'RouterParent': RouterParent,
    
    "Home": () => import('@/views/home/Home.vue'),
    'Personal': () => import('@/views/personal/index.vue'),
    // machine
    "MachineList": () => import('@/views/ops/machine/MachineList.vue'),
    "AuthCertList": () => import('@/views/ops/machine/authcert/AuthCertList.vue'),
    // sys
    "ResourceList": () => import('@/views/system/resource'),
    "RoleList": () => import('@/views/system/role'),
    "AccountList": () => import('@/views/system/account'),
    "SyslogList": () => import('@/views/system/syslog/SyslogList.vue'),
    "ConfigList": () => import('@/views/system/config/ConfigList.vue'),
    // tag
    "TagTreeList": () => import('@/views/ops/tag/TagTreeList.vue'),
    "TeamList": () => import('@/views/ops/tag/TeamList.vue'),
    // db
    "DbList": () => import('@/views/ops/db/DbList.vue'),
    "SqlExec": () => import('@/views/ops/db'),
    // redis
    "RedisList": () => import('@/views/ops/redis'),
    "DataOperation": () => import('@/views/ops/redis/DataOperation.vue'),
    // mongo
    "MongoDataOp": () => import('@/views/ops/mongo/MongoDataOp.vue'),
    // redis
    "MongoList": () => import('@/views/ops/mongo/MongoList.vue'),
}