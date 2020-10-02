package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets0b56102dad699c3ec8225ae27faa69f353659d76 = "package main\n\nimport (\n\t\"flag\"\n\n\t\"github.com/sirupsen/logrus\"\n\n\t// TODO: Fixme\n\t\"{{ .fullname }}/registry\"\n\t\"{{ .fullname }}/server\"\n)\n\nvar (\n\taddr string\n)\n\nfunc main() {\n\tflag.StringVar(&addr, \"a\", \":8080\", \"address to use\")\n\n\tcontainer, err := registry.BuildContainer()\n\tif err != nil {\n\t\tlogrus.Fatalln(err)\n\t}\n\n\terr = container.Invoke(func(s *server.Server) {\n\t\terr := s.Run(addr)\n\t\tif err != nil {\n\t\t\tlogrus.Fatalln(err)\n\t\t}\n\t})\n\tif err != nil {\n\t\tlogrus.Fatalln(err)\n\t}\n}\n"
var _Assets7d5e2066e248093e09c24d8717d50eef3a9a6af8 = "package server\n\nimport (\n\t\"github.com/gin-gonic/gin\"\n\n\t\"{{ .fullname }}/config\"\n)\n\nfunc setRoutes(engine *gin.Engine) {\n\tif config.IsProd() {\n\t\tengine.GET(\"/webpack/*name\", func(c *gin.Context) {\n\t\t\tc.File(\"server/assets/public/webpack/\" + c.Param(\"name\"))\n\t\t})\n\t}\n\tengine.GET(\"/\", func(c *gin.Context) {\n\t\tc.HTML(200, \"index.html.tmpl\", gin.H{})\n\t})\n}\n"
var _Assetseabab95ebc630ee464f076a67d6e3befd9100fd3 = "{{ .goVersion }}\n"
var _Assets70f1f0a140934f82b529769d192053de0c89868d = "# Binaries for programs and plugins\n*.exe\n*.exe~\n*.dll\n*.so\n*.dylib\n\n# Test binary, built with `go test -c`\n*.test\n\n# Output of the go coverage tool, specifically when used with LiteIDE\n*.out\n\n# Dependency directories (remove the comment below to include it)\n# vendor/\n\n# IDE settings\n.idea/\n\n# Output binary\nmain\ngin-bin\n\n# assets\nserver/assets/node_modules\nserver/assets/public\n"
var _Assets30035ce85af7cefc67ddc90da859c12f486dc558 = "// +build tools\n\npackage tools\n\nimport _ \"github.com/codegangsta/gin\"\n"
var _Assetsfffade39b597c4baa47be0c5f4fe2629a70cde93 = ".PHONY: setup\nsetup: mod-tidy npm-install build-webpack bootstrap\n\n.PHONY: mod-tidy\nmod-tidy:\n\tgo mod tidy\n\n.PHONY: npm-install\nnpm-install:\n\tcd server/assets && \\\n\tnpm install\n\n.PHONY: bootstrap\nbootstrap:\n\tgo get github.com/codegangsta/gin\n\n.PHONY: dev-server\ndev-server:\n\tgin -i -a 8080 --all -x server/assets -d cmd/server\n\n.PHONY: webpack-dev-server\nwebpack-dev-server:\n\tcd server/assets && \\\n\tnpm run dev\n\n.PHONY: build\nbuild:\n\tgo build cmd/server/*.go\n\n.PHONY: build-webpack\nbuild-webpack:\n\tcd server/assets && \\\n\tnpm run build\n"
var _Assets481a5b214de0ae8f0a166b7db8768a3aebe4b009 = "package db\n\nimport (\n\t\"github.com/pkg/errors\"\n\t\"gorm.io/driver/mysql\"\n\t\"gorm.io/gorm\"\n\n\t// TODO: fixme\n\t\"{{ .fullname }}/config\"\n)\n\nfunc ConnectDB(conf *config.AppConfig) (*gorm.DB, error) {\n\tdb, err := gorm.Open(\n\t\tmysql.New(mysql.Config{\n\t\t\tDSN: conf.DBUrl,\n\t\t}),\n\t\t&gorm.Config{\n\t\t\tAllowGlobalUpdate: false,\n\t\t})\n\tif err != nil {\n\t\treturn nil, errors.Wrap(err,\"failed to connecting database\")\n\t}\n\n\treturn db, nil\n}\n"
var _Assetsa83514d7126fe63496dc9d8a774232d99b056ce1 = "{{ .nodeVersion }}\n"
var _Assets8368c32602f94b2c8aeb83a4020f86302d43eb9c = "development:\n  database_url: root@tcp(127.0.0.1)/{{ .name }}_development?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Local\n\ntest:\n  database_url: root@tcp(127.0.0.1)/{{ .name }}_test?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Local\n\nproduction:\n  url:\n"
var _Assetsdf5a3ec2916d693a663b582aac8cf3d28d66e200 = "package registry\n\nimport (\n    \"go.uber.org/dig\"\n\n    // TODO: Fixme\n    \"{{ .fullname }}/config\"\n    \"{{ .fullname }}/db\"\n    \"{{ .fullname }}/server\"\n)\n\nfunc BuildContainer() (*dig.Container, error) {\n    c := dig.New()\n\n    providers := []*provider{\n        newProvider(config.GetConfig),\n        newProvider(db.ConnectDB),\n        newProvider(server.NewServer),\n    }\n\n    if err := setProviders(c, providers); err != nil {\n        return nil, err\n    }\n\n    return c, nil\n}\n\ntype provider struct {\n    target interface{}\n    opts   []dig.ProvideOption\n}\n\nfunc newProvider(target interface{}, opts ...dig.ProvideOption) *provider {\n    return &provider{target: target, opts: opts}\n}\n\nfunc setProviders(container *dig.Container, providers []*provider) error {\n    for _, p := range providers {\n        if err := container.Provide(p.target, p.opts...); err != nil {\n            return err\n        }\n    }\n    return nil\n}\n"
var _Assetsb5a5f7f67e0e92802dc36082dc74ba6d440f14b9 = "package main\n\nimport (\n\t\"github.com/sirupsen/logrus\"\n\t\"gorm.io/gorm\"\n\n\t\"{{ .fullname }}/registry\"\n)\n\nfunc main() {\n\tcontainer, err := registry.BuildContainer()\n\tif err != nil {\n\t\tlogrus.Fatalln(err)\n\t}\n\n\terr = container.Invoke(func(db *gorm.DB) {\n\t\t// Add models to migrate\n\t\terr := db.AutoMigrate()\n\t\tif err != nil {\n\t\t\tlogrus.Fatalln(err)\n\t\t}\n\t})\n\tif err != nil {\n\t\tlogrus.Fatalln(err)\n\t}\n}\n"
var _Assets3694895fac5b51b5cb3cf21c308ccf3231a6e19e = "{\n  \"name\": \"{{ .name }}\",\n  \"version\": \"1.0.0\",\n  \"description\": \"\",\n  \"main\": \"index.js\",\n  \"scripts\": {\n    \"dev\": \"webpack-dev-server\",\n    \"build\": \"webpack\"\n  },\n  \"devDependencies\": {\n    \"autoprefixer\": \"^10.0.1\",\n    \"compression-webpack-plugin\": \"^6.0.2\",\n    \"extract-css-chunks-webpack-plugin\": \"^4.7.5\",\n    \"postcss\": \"^8.1.1\",\n    \"webpack\": \"^4.44.2\",\n    \"webpack-cli\": \"^3.3.12\",\n    \"webpack-dev-server\": \"^3.11.0\",\n    \"webpack-manifest-plugin\": \"^2.2.0\"\n  }\n}\n"
var _Assets50b0bb6ec56abc85bdf54d7a0dff3caf8c2b4e50 = "'use strict';\n\nvar path = require('path');\nvar webpack = require('webpack');\nvar ManifestPlugin = require('webpack-manifest-plugin');\n\nvar autoprefixer = require('autoprefixer');\nvar CompressionPlugin = require(\"compression-webpack-plugin\");\n\nvar host = process.env.HOST || 'localhost'\nvar devServerPort = 3808;\n\nvar production = process.env.NODE_ENV === 'production';\n\nconst ExtractCssChunks = require(\"extract-css-chunks-webpack-plugin\")\n\nclass CleanUpExtractCssChunks {\n  shouldPickStatChild(child) {\n    return child.name.indexOf('extract-css-chunks-webpack-plugin') !== 0;\n  }\n\n  apply(compiler) {\n    compiler.hooks.done.tap('CleanUpExtractCssChunks', (stats) => {\n      const children = stats.compilation.children;\n      if (Array.isArray(children)) {\n        // eslint-disable-next-line no-param-reassign\n        stats.compilation.children = children\n            .filter(child => this.shouldPickStatChild(child));\n      }\n    });\n  }\n}\nvar config = {\n  //stats: { children: false },\n  mode: production ? \"production\" : \"development\",\n  entry: {\n    // Sources are expected to live in $app_root/webpack\n    application: 'index.js',\n  },\n\n  module: {\n    rules: [\n      { test: /\\.es6$/, use: \"babel-loader\" },\n      { test: /\\.jsx$/, use: \"babel-loader\" },\n      //{ test: /react-select\\/src/, use: \"babel-loader\" },\n      { test: /\\.(jpe?g|png|gif)$/i, use: \"file-loader\" },\n      {\n        test: /\\.woff($|\\?)|\\.woff2($|\\?)|\\.ttf($|\\?)|\\.eot($|\\?)|\\.svg($|\\?)|\\.otf($|\\?)/,\n        //use: production ? 'file-loader' : 'url-loader'\n        use: 'file-loader'\n      },\n      {\n        test: /\\.(sass|scss|css)$/,\n        use: [\n          {\n            loader: ExtractCssChunks.loader,\n            options: {\n              hot: production ? false : true,\n              // Force reload all\n              //reloadAll: true,\n            }\n          },\n          {\n            loader: \"css-loader\",\n            options: {\n              //minimize: true,\n              sourceMap: true\n            }\n          },\n          {\n            loader: \"sass-loader\"\n          }\n        ]\n      },\n    ]\n  },\n\n  output: {\n    // Build assets directly in to public/webpack/, let webpack know\n    // that all webpacked assets start with webpack/\n\n    // must match config.webpack.output_dir\n    path: path.join(__dirname, 'public', 'webpack'),\n    publicPath: '/webpack/',\n\n    filename: production ? '[name]-[chunkhash].js' : '[name].js',\n    chunkFilename: production ? '[name]-[chunkhash].js' : '[name].js',\n  },\n\n  resolve: {\n    modules: [path.resolve(__dirname, \"src\"), path.resolve(__dirname, \"node_modules\")],\n    extensions: [\".es6\", \".jsx\", \".sass\", \".css\", \".js\"],\n    alias: {\n      '~': path.resolve(__dirname, \"src\"),\n    }\n  },\n\n  plugins: [\n    new ExtractCssChunks(\n        {\n          // Options similar to the same options in webpackOptions.output\n          // both options are optional\n          filename: production ? \"[name]-[chunkhash].css\" : \"[name].css\",\n          chunkfilename: production ? \"[name]-[id].css\" : \"[name].css\",\n        }\n    ),\n    new CleanUpExtractCssChunks(),\n    new ManifestPlugin({\n      writeToFileEmit: true,\n      //basePath: \"\",\n      publicPath: production ? \"/webpack/\" : 'http://' + host + ':' + devServerPort + '/webpack/',\n    }),\n    //new webpack.IgnorePlugin(/^\\.\\/locale$/, /moment$/),\n    new webpack.ContextReplacementPlugin(/moment[/\\\\]locale$/, /ru|en/),\n  ],\n  optimization: {\n    minimize: production,\n    splitChunks: {\n      cacheGroups: {\n        default: false,\n        vendors: {\n          test: /[\\\\/]node_modules[\\\\/].*js/,\n          priority: 1,\n          name: \"vendor\",\n          chunks: \"initial\",\n          enforce: true\n        },\n      },\n    },\n  }\n};\n\nif (production) {\n  config.plugins.push(\n      //new webpack.NoEmitOnErrorsPlugin(),\n      new webpack.DefinePlugin({ // <--key to reduce React's size\n        'process.env': { NODE_ENV: JSON.stringify('production') }\n      }),\n      new CompressionPlugin({\n        //asset: \"[path].gz\",\n        algorithm: \"gzip\",\n        test: /\\.js$|\\.css$/,\n        threshold: 4096,\n        minRatio: 0.8\n      })\n  );\n  config.output.publicPath = '/webpack/';\n} else {\n  config.plugins.push(\n      new webpack.NamedModulesPlugin(),\n  )\n\n  config.devServer = {\n    port: devServerPort,\n    disableHostCheck: true,\n    headers: { 'Access-Control-Allow-Origin': '*' },\n  };\n\n  config.output.publicPath = 'http://' + host + ':' + devServerPort + '/webpack/';\n  // Source maps\n  config.devtool = 'source-map';\n}\n\nmodule.exports = config\n"
var _Assetsf75c87b9097f8f740b104807f94aec87745cde3a = "package server\n\nimport (\n\t\"html/template\"\n\t\"path/filepath\"\n\t\"time\"\n\n\t\"github.com/gin-contrib/multitemplate\"\n\t\"github.com/gin-gonic/gin\"\n\t\"github.com/go-webpack/webpack\"\n\t\"github.com/sirupsen/logrus\"\n\tginlogrus \"github.com/toorop/gin-logrus\"\n\n\t// TODO: Fixme\n\t\"{{ .fullname }}/config\"\n)\n\nfunc init() {\n\tlogrus.SetFormatter(&logrus.TextFormatter{\n\t\tFullTimestamp:   true,\n\t\tTimestampFormat: time.RFC3339,\n\t})\n}\n\ntype Server struct {\n\t*gin.Engine\n}\n\nfunc NewServer() *Server {\n\tsetupWebpack()\n\n\tengine := gin.New()\n\tengine.Use(ginlogrus.Logger(logrus.StandardLogger()), gin.Recovery())\n\tengine.HTMLRender = loadTemplates(\"server/views\")\n\n\tsetRoutes(engine)\n\n\treturn &Server{Engine: engine}\n}\n\nfunc setupWebpack() {\n\twebpack.DevHost = \"localhost:3808\" // default\n\twebpack.Plugin = \"manifest\"        // defaults to stats for compatibility\n\twebpack.FsPath = \"./server/assets/public/webpack\"\n\twebpack.Init(config.IsDev())\n}\n\nfunc loadTemplates(templatesDir string) multitemplate.Renderer {\n\tr := multitemplate.NewRenderer()\n\n\tlayouts, err := filepath.Glob(templatesDir + \"/layouts/*.html.tmpl\")\n\tif err != nil {\n\t\tpanic(err.Error())\n\t}\n\n\tpages, err := filepath.Glob(templatesDir + \"/pages/*.html.tmpl\")\n\tif err != nil {\n\t\tpanic(err.Error())\n\t}\n\n\tfuncMap := template.FuncMap{\n\t\t\"asset\": webpack.AssetHelper,\n\t}\n\n\t// Generate our templates map from our layouts/ and pages/ directories\n\tfor _, include := range pages {\n\t\tlayoutCopy := make([]string, len(layouts))\n\t\tcopy(layoutCopy, layouts)\n\t\tfiles := append(layoutCopy, include)\n\t\tr.AddFromFilesFuncs(filepath.Base(include), funcMap, files...)\n\t}\n\treturn r\n}\n"
var _Assets47609cde872c6d91f22670c269971d1b06c8d019 = "<!DOCTYPE html>\n<html>\n<head>\n  <meta charset=\"utf-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    {{\"{{\"}} asset \"vendor.js\" {{\"}}\"}}\n    {{\"{{\"}} asset \"main.js\" {{\"}}\"}}\n  <title>{{ .name }}</title>\n</head>\n<body>\n{{\"{{\"}} template \"content\" . {{\"}}\"}}\n</body>\n</html>\n"
var _Assets3c95a58eccf1beb43dae088d952505d82b3b7e7e = "{{\"{{\"}} define \"content\" {{\"}}\"}}\nHello, {{ .name }}!\n{{ \"{{ end }}\" }}\n"
var _Assetse11834f62b5e46ac61ed2f1e5146d14ccbb9fe5c = "module {{ .fullname }}\n\ngo {{ .goMinorVersion }}\n\nrequire (\n)\n"
var _Assetsf5b4e7350c57dd57af3cb233497a017b06bd1400 = "package config\n\nimport (\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"os\"\n\n\t\"gopkg.in/yaml.v2\"\n)\n\nconst Dev = \"development\"\nconst Test = \"test\"\nconst Prod = \"production\"\n\ntype AppConfig struct {\n\tDBUrl string `yaml:\"database_url\"`\n}\n\nfunc GetConfig() (*AppConfig, error) {\n\tenv := GetEnv()\n\n\tconf, err := readSettings(env)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\treturn conf, nil\n}\n\nfunc GetEnv() string {\n\tenv := os.Getenv(\"GO_APP_ENV\")\n\tif env == \"\" {\n\t\tenv = Dev\n\t}\n\n\treturn env\n}\n\nfunc IsDev() bool {\n\treturn GetEnv() == Dev\n}\n\nfunc IsTest() bool {\n\treturn GetEnv() == Test\n}\n\nfunc IsProd() bool {\n\treturn GetEnv() == Prod\n}\n\nfunc readSettings(env string) (*AppConfig, error) {\n\tfile, err := ioutil.ReadFile(\"config/settings.yml\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tfile = []byte(os.ExpandEnv(string(file)))\n\tconfs := make(map[string]*AppConfig)\n\tif err := yaml.Unmarshal(file, confs); err != nil {\n\t\treturn nil, err\n\t}\n\n\tconf := confs[env]\n\tif conf == nil {\n\t\treturn nil, fmt.Errorf(\"environment '%s' is not found on config/settings.yml\", env)\n\t}\n\n\treturn conf, nil\n}\n"
var _Assets619c91612e0168b92064357135659f9e86c8503f = ""
var _Assetsc3b8bbeffd0f185321d8803c71c4ca87a5cfdd45 = "# {{ .name }}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"assets"}, "/assets": []string{"Makefile.tmpl", ".gitignore.tmpl", "README.md.tmpl", "go.mod.tmpl", ".go-version.tmpl", ".node-version.tmpl"}, "/assets/cmd": []string{}, "/assets/cmd/migrate": []string{"main.go.tmpl"}, "/assets/cmd/server": []string{"main.go.tmpl"}, "/assets/config": []string{"app_config.go.tmpl", "settings.yml.tmpl"}, "/assets/db": []string{"db.go.tmpl"}, "/assets/model": []string{}, "/assets/registry": []string{"container.go.tmpl"}, "/assets/server": []string{"route.go.tmpl", "server.go.tmpl"}, "/assets/server/assets": []string{"package.json.tmpl", "webpack.config.js.tmpl"}, "/assets/server/assets/src": []string{"index.js"}, "/assets/server/views": []string{}, "/assets/server/views/layouts": []string{"base.html.tmpl.tmpl"}, "/assets/server/views/pages": []string{"index.html.tmpl.tmpl"}, "/assets/tools": []string{"tools.go.tmpl"}}, map[string]*assets.File{
	"/assets/config/app_config.go.tmpl": &assets.File{
		Path:     "/assets/config/app_config.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601739792, 1601739792563887481),
		Data:     []byte(_Assetsf5b4e7350c57dd57af3cb233497a017b06bd1400),
	}, "/assets/server/assets/src": &assets.File{
		Path:     "/assets/server/assets/src",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601759572, 1601759572926664903),
		Data:     nil,
	}, "/assets/server/assets/src/index.js": &assets.File{
		Path:     "/assets/server/assets/src/index.js",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601759572, 1601759572926651069),
		Data:     []byte(_Assets619c91612e0168b92064357135659f9e86c8503f),
	}, "/assets/README.md.tmpl": &assets.File{
		Path:     "/assets/README.md.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601582986, 1601582986211481571),
		Data:     []byte(_Assetsc3b8bbeffd0f185321d8803c71c4ca87a5cfdd45),
	}, "/assets/go.mod.tmpl": &assets.File{
		Path:     "/assets/go.mod.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601757174, 1601757174747477929),
		Data:     []byte(_Assetse11834f62b5e46ac61ed2f1e5146d14ccbb9fe5c),
	}, "/assets/cmd": &assets.File{
		Path:     "/assets/cmd",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601742873, 1601742873191968604),
		Data:     nil,
	}, "/assets/server/route.go.tmpl": &assets.File{
		Path:     "/assets/server/route.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601759119, 1601759119227826581),
		Data:     []byte(_Assets7d5e2066e248093e09c24d8717d50eef3a9a6af8),
	}, "/assets/server/views": &assets.File{
		Path:     "/assets/server/views",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601740576, 1601740576249637762),
		Data:     nil,
	}, "/assets/.go-version.tmpl": &assets.File{
		Path:     "/assets/.go-version.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601657865, 1601657865797246403),
		Data:     []byte(_Assetseabab95ebc630ee464f076a67d6e3befd9100fd3),
	}, "/assets/db": &assets.File{
		Path:     "/assets/db",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601756937, 1601756937832565051),
		Data:     nil,
	}, "/assets/cmd/server/main.go.tmpl": &assets.File{
		Path:     "/assets/cmd/server/main.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601756937, 1601756937829319964),
		Data:     []byte(_Assets0b56102dad699c3ec8225ae27faa69f353659d76),
	}, "/assets/.gitignore.tmpl": &assets.File{
		Path:     "/assets/.gitignore.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601761554, 1601761554877517982),
		Data:     []byte(_Assets70f1f0a140934f82b529769d192053de0c89868d),
	}, "/assets/cmd/server": &assets.File{
		Path:     "/assets/cmd/server",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601756937, 1601756937829405376),
		Data:     nil,
	}, "/assets/tools/tools.go.tmpl": &assets.File{
		Path:     "/assets/tools/tools.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601761247, 1601761247419853963),
		Data:     []byte(_Assets30035ce85af7cefc67ddc90da859c12f486dc558),
	}, "/assets/Makefile.tmpl": &assets.File{
		Path:     "/assets/Makefile.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601761542, 1601761542937968470),
		Data:     []byte(_Assetsfffade39b597c4baa47be0c5f4fe2629a70cde93),
	}, "/assets/db/db.go.tmpl": &assets.File{
		Path:     "/assets/db/db.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601756937, 1601756937832431614),
		Data:     []byte(_Assets481a5b214de0ae8f0a166b7db8768a3aebe4b009),
	}, "/assets/.node-version.tmpl": &assets.File{
		Path:     "/assets/.node-version.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601755679, 1601755679353991681),
		Data:     []byte(_Assetsa83514d7126fe63496dc9d8a774232d99b056ce1),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601761353, 1601761353497711809),
		Data:     nil,
	}, "/assets/server/views/pages": &assets.File{
		Path:     "/assets/server/views/pages",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601759405, 1601759405307667104),
		Data:     nil,
	}, "/assets/config/settings.yml.tmpl": &assets.File{
		Path:     "/assets/config/settings.yml.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601742049, 1601742049559700120),
		Data:     []byte(_Assets8368c32602f94b2c8aeb83a4020f86302d43eb9c),
	}, "/assets/tools": &assets.File{
		Path:     "/assets/tools",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601761262, 1601761262348801985),
		Data:     nil,
	}, "/assets/server": &assets.File{
		Path:     "/assets/server",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601760404, 1601760404209338702),
		Data:     nil,
	}, "/assets/server/views/layouts": &assets.File{
		Path:     "/assets/server/views/layouts",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601741355, 1601741355730290208),
		Data:     nil,
	}, "/assets/server/assets/package.json.tmpl": &assets.File{
		Path:     "/assets/server/assets/package.json.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601756788, 1601756788280349923),
		Data:     []byte(_Assets3694895fac5b51b5cb3cf21c308ccf3231a6e19e),
	}, "/assets/registry/container.go.tmpl": &assets.File{
		Path:     "/assets/registry/container.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601756937, 1601756937822542815),
		Data:     []byte(_Assetsdf5a3ec2916d693a663b582aac8cf3d28d66e200),
	}, "/assets/cmd/migrate/main.go.tmpl": &assets.File{
		Path:     "/assets/cmd/migrate/main.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601743052, 1601743052822915928),
		Data:     []byte(_Assetsb5a5f7f67e0e92802dc36082dc74ba6d440f14b9),
	}, "/assets/config": &assets.File{
		Path:     "/assets/config",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601756915, 1601756915928711075),
		Data:     nil,
	}, "/assets/server/server.go.tmpl": &assets.File{
		Path:     "/assets/server/server.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601760287, 1601760287236837196),
		Data:     []byte(_Assetsf75c87b9097f8f740b104807f94aec87745cde3a),
	}, "/assets/server/views/layouts/base.html.tmpl.tmpl": &assets.File{
		Path:     "/assets/server/views/layouts/base.html.tmpl.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601741355, 1601741355730165255),
		Data:     []byte(_Assets47609cde872c6d91f22670c269971d1b06c8d019),
	}, "/assets/server/assets": &assets.File{
		Path:     "/assets/server/assets",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601759539, 1601759539183425205),
		Data:     nil,
	}, "/assets/server/assets/webpack.config.js.tmpl": &assets.File{
		Path:     "/assets/server/assets/webpack.config.js.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601756707, 1601756707171895613),
		Data:     []byte(_Assets50b0bb6ec56abc85bdf54d7a0dff3caf8c2b4e50),
	}, "/assets/cmd/migrate": &assets.File{
		Path:     "/assets/cmd/migrate",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601743052, 1601743052822992227),
		Data:     nil,
	}, "/assets/server/views/pages/index.html.tmpl.tmpl": &assets.File{
		Path:     "/assets/server/views/pages/index.html.tmpl.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1601759405, 1601759405307575943),
		Data:     []byte(_Assets3c95a58eccf1beb43dae088d952505d82b3b7e7e),
	}, "/assets/model": &assets.File{
		Path:     "/assets/model",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601742958, 1601742958261776582),
		Data:     nil,
	}, "/assets/registry": &assets.File{
		Path:     "/assets/registry",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601756937, 1601756937822676284),
		Data:     nil,
	}, "/assets": &assets.File{
		Path:     "/assets",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1601761554, 1601761554877650887),
		Data:     nil,
	}}, "")
