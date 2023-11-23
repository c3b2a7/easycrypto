<script setup>
import { Decrypt, Encrypt } from 'wailsjs/go/service/CryptoService.js'
import { ElNotification } from 'element-plus'
import { ClipboardSetText } from 'wailsjs/runtime/runtime.js'

const options = [
    { label: '萨摩耶应用数据库', value: 'smy' },
    { label: '浏览器数据库', value: 'browser-db' },
    { label: '浏览器测试环境接口', value: 'browser-dev' },
    { label: '浏览器生产环境接口', value: 'browser-prd' },
]

const state = reactive({
    plaintext: '',
    ciphertext: '',
    env: 'smy',
})

const autoClipboard = ref(false)

const encrypt = () => {
    Encrypt({ env: state.env, data: state.plaintext }).then((result) => {
        if (result.success) {
            if (autoClipboard.value) {
                ClipboardSetText(result.data)
            }
            state.ciphertext = result.data
        } else {
            showErrorNotification('加密失败', result.msg)
        }
    })
}

const decrypt = () => {
    Decrypt({ env: state.env, data: state.ciphertext }).then((result) => {
        if (result.success) {
            if (autoClipboard.value) {
                ClipboardSetText(result.data)
            }
            state.plaintext = result.data
        } else {
            showErrorNotification('解密失败', result.msg)
        }
    })
}

const showErrorNotification = (title, message) => {
    let n = ElNotification({
        title,
        message,
        type: 'error',
        duration: 3000,
        showClose: false,
        onClick: () => n.close(),
    })
}
</script>

<template>
    <el-container>
        <el-header><p>加解密小工具</p></el-header>
        <el-main>
            <el-form :model="state">
                <el-row>
                    <el-col>
                        <el-form-item>
                            <el-input
                                v-model="state.plaintext"
                                type="textarea"
                                placeholder="需加密的内容粘贴在这里"
                                resize="none"
                                rows="5" />
                        </el-form-item>
                    </el-col>
                    <el-col>
                        <el-form-item>
                            <el-input
                                v-model="state.ciphertext"
                                type="textarea"
                                placeholder="需解密的内容粘贴在这里"
                                resize="none"
                                rows="5" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="6">
                        <el-form-item>
                            <el-select v-model="state.env">
                                <el-option
                                    v-for="(item, index) in options"
                                    :key="index"
                                    :label="item.label"
                                    :value="item.value" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="18">
                        <el-form-item style="float: right">
                            <el-checkbox v-model="autoClipboard" label="自动复制结果到剪贴板" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col>
                        <el-form-item style="float: right">
                            <el-button type="primary" @click="encrypt">加密</el-button>
                            <el-button type="primary" @click="decrypt">解密</el-button>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
        </el-main>
    </el-container>
</template>

<style scoped></style>