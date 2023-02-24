module.exports = class Utility {
    genRandonString(length) {
        const chars = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()';
        const charLength = chars.length;
        let result = '';
        for ( var i = 0; i < length; i++ ) {
            result += chars.charAt(Math.floor(Math.random() * charLength));
        }
        return result;
    }
    median(values) {
        if (values.length === 0) {
            return 0;
        }

        values.sort(function(a, b) {
            return a - b;
        });

        const half = Math.floor(values.length / 2);

        if (values.length % 2) {
            return values[half];
        } else {
            return (values[half - 1] + values[half]) / 2.0;
        }
    }

    average(values) {
        if (values.length === 0) {
            return 0;
        }

        const sum = values.reduce((acc, curr) => acc + curr);
        return sum / values.length;
    }
}
