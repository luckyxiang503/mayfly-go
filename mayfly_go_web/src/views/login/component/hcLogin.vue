<template>
    <el-form ref="loginFormRef" :model="loginForm" :rules="rules" class="login-content-form" size="large">
        <el-form-item prop="username">
            <el-input type="text" placeholder="请输入用户名" prefix-icon="el-icon-user" v-model="loginForm.username" clearable autocomplete="off">
            </el-input>
        </el-form-item>
        <el-form-item prop="code">
            <el-row :gutter="15">
                <el-col :span="16">
                    <el-input
                        type="text"
                        maxlength="10"
                        placeholder="请输入验证码"
                        prefix-icon="el-icon-position"
                        v-model="loginForm.code"
                        clearable
                        autocomplete="off"
                        @keyup.enter="login"
                    ></el-input>
                </el-col>
                <el-col :span="8">
                    <el-button @click="getCode">获取验证码</el-button>
                </el-col>
            </el-row>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" class="login-content-submit" round @click="login">
                <span>登 录</span>
            </el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import { ref, toRefs, reactive, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { initBackEndControlRoutesFun } from '@/router/index.ts';
import { useStore } from '@/store/index.ts';
import { setSession, setUserInfo2Session, setUseWatermark2Session } from '@/common/utils/storage.ts';
import { formatAxis } from '@/common/utils/formatTime.ts';
import openApi from '@/common/openApi';
import {useWartermark } from '@/common/sysconfig';
import { letterAvatar } from '@/common/utils/string';

const rules = {
    username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
    code: [{ min: 6, max: 6, message: '验证码为6位字符', trigger: 'blur' }],
}

const store = useStore();
const route = useRoute();
const router = useRouter();
const loginFormRef: any = ref(null);

const state = reactive({
    loginForm: {
        username: '',
        code: '',
    },
    loading: {
        signIn: false,
    },
});

const {
    loginForm,
    loading
} = toRefs(state)


// 时间获取
const currentTime = computed(() => {
    return formatAxis(new Date());
});

// 调用接口让后台发送验证码
const getCode = () => {
    loginFormRef.value.validateField('username', (valid: boolean) => {
        if (valid) {
            onSendCode();
        } else {
            return false;
        }
    });
}

// 发送验证码
const onSendCode =async () => {
    try {
        const loginReq = { username: state.loginForm.username };
        await openApi.sendCode(loginReq);
        ElMessage.success("验证码已发送，有效时间为120s")
    } catch (e: any) {
        console.log(e)
        return;
    }
};

// 校验登录表单并登录
const login = () => {
    loginFormRef.value.validate((valid: boolean) => {
        if (valid) {
            onSignIn();
        } else {
            return false;
        }
    });
};

// 登录
const onSignIn = async () => {
    state.loading.signIn = true;
    let loginRes;
    try {
        const loginReq = { ...state.loginForm };
        loginRes = await openApi.hclogin(loginReq);
        // 存储 token 到浏览器缓存
        setSession('token', loginRes.token);
        setSession('menus', loginRes.menus);
    } catch (e: any) {
        state.loading.signIn = false;
        console.log(e)
        return;
    }
    // 用户信息
    const userInfos = {
        name: loginRes.name,
        username: state.loginForm.username,
        // 头像
        photo: letterAvatar(state.loginForm.username),
        time: new Date().getTime(),
        // // 菜单资源code数组
        // menus: loginRes.menus,
        permissions: loginRes.permissions,
        lastLoginTime: loginRes.lastLoginTime,
        lastLoginIp: loginRes.lastLoginIp,
    };

    // 存储用户信息到浏览器缓存
    setUserInfo2Session(userInfos);
    // 1、请注意执行顺序(存储用户信息到vuex)
    store.dispatch('userInfos/setUserInfos', userInfos);
    if (!store.state.themeConfig.themeConfig.isRequestRoutes) {
        // 前端控制路由，2、请注意执行顺序
        // await initAllFun();
        await initBackEndControlRoutesFun();
        signInSuccess();
    } else {
        // 模拟后端控制路由，isRequestRoutes 为 true，则开启后端控制路由
        // 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
        await initBackEndControlRoutesFun();
        // 执行完 initBackEndControlRoutesFun，再执行 signInSuccess
        signInSuccess();
    }
};

// 登录成功后的跳转
const signInSuccess = () => {
    // 初始化登录成功时间问候语
    let currentTimeInfo = currentTime.value;
    // 登录成功，跳到转首页
    // 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
    // 如果是复制粘贴的路径，非首页/登录页，那么登录成功后重定向到对应的路径中
    route.query?.redirect ? router.push(route.query.redirect as string) : router.push('/');
    // 登录成功提示
    setTimeout(async () => {
        // 关闭 loading
        state.loading.signIn = true;
        ElMessage.success(`${currentTimeInfo}，欢迎回来！`);
        if (await useWartermark()) {
            setUseWatermark2Session(true);
        }
    }, 300);
};
</script>

<style scoped lang="scss">
.login-content-form {
    margin-top: 20px;
    .login-content-submit {
        width: 100%;
        letter-spacing: 2px;
        font-weight: 300;
        margin-top: 15px;
    }
}
</style>
