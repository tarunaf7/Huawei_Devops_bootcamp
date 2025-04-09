require('@google-cloud/trace-agent').start();

const path = require('path');
const grpc = require('grpc');
const pino = require('pino');

const PROTO_PATH = path.join(__dirname, './proto/demo.proto');
const PORT = 7000;

const shopProto = grpc.load(PROTO_PATH).hipstershop;
const client = new shopProto.CurrencyService(`localhost:${PORT}`,
  grpc.credentials.createInsecure());

const logger = pino({
  name: 'currencyservice-client',
  messageKey: 'message',
  formatters: {
    level (logLevelString, logLevelNum) {
      return { severity: logLevelString }
    }
  }
});

const request = {
  from: {
    currency_code: 'CHF',
    units: 300,
    nanos: 0
  },
  to_code: 'EUR'
};

function _moneyToString (m) {
  return `${m.units}.${m.nanos.toString().padStart(9,'0')} ${m.currency_code}`;
}

client.getSupportedCurrencies({}, (err, response) => {
  if (err) {
    logger.error(`Error in getSupportedCurrencies: ${err}`);
  } else {
    logger.info(`Currency codes: ${response.currency_codes}`);
  }
});

client.convert(request, (err, response) => {
  if (err) {
    logger.error(`Error in convert: ${err}`);
  } else {
    logger.log(`Convert: ${_moneyToString(request.from)} to ${_moneyToString(response)}`);
  }
});
