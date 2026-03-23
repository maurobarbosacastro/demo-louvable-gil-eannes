const proxy = {
    '/awin': (req, res) => {
        return res.json(require('./awin.json'));
    },
}
module.exports = proxy;
