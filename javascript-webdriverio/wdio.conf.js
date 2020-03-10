exports.config = {
    runner: 'local',
    protocol: 'http',
    hostname: '127.0.0.1',
    port: 4444,
    //user: '', // Username
    //key: '', // Password
    specs: [
        'tests/**/*js'
    ],
    maxInstances: 1,
    capabilities: [
      {browserName: 'chrome', browserVersion: '79.0',
        "selenoid:options": { "enableVNC": true }
      }
    ],
    logLevel: 'debug',
    framework: 'mocha',
    reporters: ['dot','spec'],
    mochaOpts: {
        ui: 'bdd',
        timeout: 60000
    }
};
