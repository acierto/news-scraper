module.exports = function(config) {
    var browser = process.env.KARMA_BROWSER || 'PhantomJS';

    config.set({
        basePath: process.env.PROJECT_DIR + '/web',
        frameworks: ['jasmine'],
        files: [
            'libs/js/lodash/lodash.compat.min.js',
            'libs/js/angular/angular.min.js',
            'libs/js/angular-resource/angular-resource.min.js',
            'libs/js/angular-route/angular-route.min.js',
            'libs/js/angular-sanitize/angular-sanitize.min.js',
            'libs/js/angular-mocks/angular-mocks.js',
            'libs/js/jquery/jquery.js',
            'js/*.js',
            'js/services/*.js',
            'js/filters/*.js',
//            'js/directives/*.js',
            'js/controllers/*.js',
            '../test/javascript/unit/**/*.coffee',
            '**/*.html'
        ],
        preprocessors: {
            'partials/directives/**/*.html': 'ng-html2js',
            '../../../**/*.coffee': 'coffee',
            '**/*.html' : ["ng-html2js"]
        },
        ngHtml2JsPreprocessor: {
            moduleName: "templates"
        },
        junitReporter: {
            outputFile: process.env.PROJECT_DIR + '/build/test-results/karma-test-results.xml'
        },
        exclude: [],
        reporters: ['progress', 'coverage', 'junit'],
        port: 9997,
        runnerPort: 9100,
        colors: true,
        logLevel: config.LOG_INFO,
        autoWatch: false,
        browsers: [browser],
        captureTimeout: 10000,
        singleRun: true
    });
};