<script setup>
import { Decrypt, Encrypt, GetKeys } from 'wailsjs/go/service/CryptoService.js'
import { ElNotification } from 'element-plus'
import { ClipboardSetText } from 'wailsjs/runtime/runtime.js'

const options = ref([])

const state = reactive({
    plaintext: '',
    ciphertext: '',
    env: '',
})

const autoClipboard = ref(false)

onMounted(() => {
    GetKeys().then((keys) => {
        if (keys && keys.length >= 1) {
            keys.forEach((key) => {
                options.value.push({ label: key, value: key })
            })
            state.env = keys[0]
        }
    })
})

const encrypt = () => {
    Encrypt(state.env, state.plaintext)
        .then((result) => {
            if (autoClipboard.value) {
                ClipboardSetText(result)
            }
            state.ciphertext = result
        })
        .catch((error) => {
            showErrorNotification('加密失败', error)
        })
}

const decrypt = () => {
    Decrypt(state.env, state.ciphertext)
        .then((result) => {
            if (autoClipboard.value) {
                ClipboardSetText(result)
            }
            state.plaintext = result
        })
        .catch((error) => {
            showErrorNotification('解密失败', error)
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
