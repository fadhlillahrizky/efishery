const db = require("../../config/db");
const model = require("../../models/index");

module.exports = class service {
    constructor(option) {
        this.db = option?.db || db;
        this.model = option?.model || model;
    }

    async findUser(phone, password ) {
        return this.model.Users.findOne({
            where: {
                phone, password
            }
        })
    }

    async findByPhone(phone) {
        return this.model.Users.findOne({
            where: {
                phone
            }
        })
    }

    async createUser(body) {
        return this.model.Users.create(body)
    }
}
