const GenericResponseEntity = require("../../utilities/entities");
const jwt = require("jsonwebtoken");
const Repository = require("./repository");
const Utility = require("../../utilities/utility");
const jwtSecret = process.env.SECRET ?? "secretkey";

module.exports = class service {
    constructor(option) {
        this.repository = option?.repository || new Repository();
        this.utility = option?.utility || new Utility();
    }

    async login(body) {
        const response = new GenericResponseEntity();
        const { phone, password } = body;

        const user = await this.repository.findUser(phone, password);
        if (!user) {
            response.statusCode = 400;
            response.message = "Invalid phone or password";
            return response;
        }

        const token = jwt.sign({
            id: user.id,
            phone: user.phone,
            name: user.name,
            role: user.role,
        }, jwtSecret);

        response.message = "success";
        response.success = true;
        response.data = {
            token,
        };
        return response;
    }

    async register(body) {
        const response = new GenericResponseEntity();
        const { phone, name, role } = body;

        const user = await this.repository.findByPhone(phone);
        if (user) {
            response.statusCode = 400;
            response.message = "Phone already used";
            return response;
        }
        const password = this.utility.genRandonString(4);
        await this.repository.createUser({
            phone,
            name,
            role,
            password,
        });

        response.message = "success";
        response.success = true;
        response.data = {
            phone,
            name,
            role,
            password,
        };
        return response;
    }

    async checkToken(token) {
        const response = new GenericResponseEntity();

        try {
            const decoded = jwt.verify(token, jwtSecret);

            response.message = "success";
            response.success = true;
            response.data = decoded;
            return response;
        } catch (e) {
            response.statusCode = 400;
            response.message = "Invalid Token";
            return response;
        }
    }
}
