const GenericResponseEntity = require("../utilities/entities");
const { httpResponse } = require("../utilities/response");
const jwt = require("jsonwebtoken");
const jwtSecret = process.env.SECRET ?? "secretkey";

module.exports = (req, res, next) => {

    const response = new GenericResponseEntity();
    response.statusCode = 401;
    response.message = "Auth Failed";
    const token = req.header("authorization").split(' ')[1] ?? '';
    if (!token) {
        response.statusCode = 401;
        response.message = "Access Denied";
        return httpResponse(response, res);
    }
    try {
        const decoded = jwt.verify(token, jwtSecret);
        req.user = decoded;
        next();
    } catch (err) {
        response.statusCode = 400;
        response.message = "Invalid Token";
        return httpResponse(response, res);
    }
};
