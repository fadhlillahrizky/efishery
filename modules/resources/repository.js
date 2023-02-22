const axios = require("axios");
const baseUrl = process.env.EFISHERY_BASE_URL ?? "https://stein.efishery.com/v1/";

module.exports = class service {
    constructor(option) {
        this.axios = option?.axios || axios;
        this.baseUrl = option?.baseUrl || baseUrl;
    }

    async list() {
        let url = `${this.baseUrl}storages/5e1edf521073e315924ceab4/list`;
        return await this.axios.get(url);
    }

    async get(id) {
        let url = `${this.baseUrl}recruitment/positions/${id}`;
        return await this.axios.get(url);
    }
}
