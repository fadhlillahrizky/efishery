const axios = require("axios");
const NodeCache = require("node-cache");

module.exports = class service {
    constructor(option) {
        this.axios = option?.axios || axios;
        this.baseUrl = option?.baseUrl || process.env.EFISHERY_BASE_URL;
        this.baseCurrencyUrl = option?.baseCurrencyUrl || process.env.CURRENCY_BASE_URL;
        // this.cache = option?.cache || new NodeCache({ stdTTL: 15 });
    }

    async list() {
        let url = `${this.baseUrl}storages/5e1edf521073e315924ceab4/list`;
        return await this.axios.get(url);
    }

    async getCurrency() {
        const apiKey = process.env.CURRENCY_API_KEY;
        let url = `${this.baseCurrencyUrl}?apikey=${apiKey}&currencies=IDR`;
        return await this.axios.get(url);
    }
}
