const GenericResponseEntity = require("../../utilities/entities");
const Repository = require("./repository");
const dayjs = require("dayjs");
const Utility = require("../../utilities/utility");

module.exports = class service {
    constructor(option) {
        this.repository = option?.repository || new Repository();
        this.utility = option?.utility || new Utility();
    }

    async list(nodeCache) {
        const response = new GenericResponseEntity();

        let usdCurrency =  nodeCache.get('usd');
        if (!nodeCache.has('usd')) {
            usdCurrency = (await this.repository.getCurrency())?.data.data?.IDR ?? 1;
            nodeCache.set('usd', usdCurrency)
        }
        try {
            const resources = (await this.repository.list());

            response.message = "success";
            response.success = true;
            response.data = resources.data.map((resource) => {
                resource.price_usd = (parseInt(resource.price) / usdCurrency).toFixed(2);
                return resource;
            });
            return response;
        } catch (e) {
            response.statusCode = 500;
            response.message = "Error Server";
            response.errors = e;
            return response;
        }

    }

    async aggregate() {
        const response = new GenericResponseEntity();

        try {
            const resources = await this.repository.list();
            const aggregatedData = {};

            resources.data.forEach((item) => {
                const area = item.area_provinsi;
                const date = dayjs(item.tgl_parsed).format("YYYY-MM-DD");

                if (!aggregatedData[area]) {
                    aggregatedData[area] = {};
                }

                if (!aggregatedData[area][date]) {
                    aggregatedData[area][date] = {
                        prices: [],
                        sizes: []
                    };
                }

                aggregatedData[area][date].prices.push(item.price);
                aggregatedData[area][date].sizes.push(item.size);

            });

            const result = {};
            Object.keys(aggregatedData).forEach((area) => {
                result[area] = {};
                Object.keys(aggregatedData[area]).forEach((date) => {
                    const prices = aggregatedData[area][date].prices;
                    const sizes = aggregatedData[area][date].sizes;
                    result[area][date] = {
                        min_price: Math.min(...prices),
                        max_price: Math.max(...prices),
                        median_price: this.utility.median(prices),
                        avg_price: this.utility.average(prices),
                        min_size: Math.min(...sizes),
                        max_size: Math.max(...sizes),
                        median_size: this.utility.median(sizes),
                        avg_size: this.utility.average(sizes)
                    };
                });
            });

            response.message = "success";
            response.success = true;
            response.data = result;
            return response;
        } catch (e) {
            response.statusCode = 500;
            response.message = "Error Server";
            response.data = e;
            return response;
        }
    }
}
