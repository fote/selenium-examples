exports.config = {

  seleniumAddress: `http://127.0.0.1:4444/wd/hub`,
  maxSessions: 2,

  capabilities: {
    'browserName': 'chrome'
  },

  specs: [
    'example_spec.js'
  ]

};
