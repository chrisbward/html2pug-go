package pkg_test

import (
	"testing"

	pkg "github.com/chrisbward/html2pug-go/pkg/html2pug-go"
	"github.com/chrisbward/html2pug-go/pkg/html2pug-go/entities"
	"github.com/sirupsen/logrus"
	assert "github.com/stretchr/testify/assert"
	_ "go.uber.org/mock/gomock"
)

func TestConvert(t *testing.T) {

	defaultOptions := &entities.Html2JadeConvertorOptions{
		NSpaces: 2,
	}
	defaultOptionsWithHead := &entities.Html2JadeConvertorOptions{
		NSpaces:  2,
		KeepHead: true,
	}
	doSKip := true

	type TestCase struct {
		Desc         string
		Skip         bool
		Options      *entities.Html2JadeConvertorOptions
		SourceHTML   string
		ExpectedJade string
		NilAssertion func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}

	testCases := []TestCase{
		{
			Desc:    "TEST000 - Partial",
			Options: defaultOptions,
			SourceHTML: `<p>hello world</p>
  `,
			ExpectedJade: `html
  body
    p hello world
`,
			NilAssertion: assert.NotNil,
		},
		{
			Desc:    "TEST001 - Evaluate Angular",
			Options: defaultOptions,
			SourceHTML: `<button
  ng-click="login('testing', 'testing')"
  class="btn btn-small">Test Login</button>
`,
			ExpectedJade: `html
  body
    button.btn.btn-small(ng-click="login('testing', 'testing')") Test Login
`,
			NilAssertion: assert.NotNil,
		},
		{
			Desc:    "TEST002 - Apostrophes",
			Options: defaultOptions,
			SourceHTML: `<img title="Joe's Place" />
<img title='Joe"s Place' />
`,
			ExpectedJade: `html
  body
    img(title="Joe's Place")
    img(title='Joe"s Place')
`,
			NilAssertion: assert.Nil,
		},
		{
			Desc:    "TEST003 - Attributes, multi-line",
			Options: defaultOptions,
			SourceHTML: `<img src="img/close_button.png" height="16" width="16" alt="Home"
onclick="
    mwl.switchClass('#search_title', 'show_title_search', 'show_title_main');
    mwl.setGroupTarget('#navigateToggle', '#home', 'show', 'hide');
    mwl.switchClass('#slider', 'show_miniapp', 'show_main');
    mwl.scrollTo('#main');"/>
`,
			ExpectedJade: `html
  body
    img(src='img/close_button.png', height='16', width='16', alt='Home', onclick="\
    mwl.switchClass('#search_title', 'show_title_search', 'show_title_main');\
    mwl.setGroupTarget('#navigateToggle', '#home', 'show', 'hide');\
    mwl.switchClass('#slider', 'show_miniapp', 'show_main');\
    mwl.scrollTo('#main');")
`,
			NilAssertion: assert.Nil,
		},
		{
			Desc:    "TEST004 - Code inline",
			Options: defaultOptions,
			SourceHTML: `<code>inline</code>
`,
			ExpectedJade: `html
  body
    code inline
`,
			NilAssertion: assert.Nil,
		}, {
			Desc:    "TEST005 - Code multi-line",
			Options: defaultOptions,
			SourceHTML: `<code>
blah
blah
</code>
`,
			ExpectedJade: `html
  body
    code
      | blah
      | blah
`,
			NilAssertion: assert.Nil,
		},
		// 		{
		// 			Desc:    "TEST006 - Comment",
		// 			Options: defaultOptions,
		// 			SourceHTML: `<html>
		//   <head>
		//   </head>
		//   <body>
		//     <!--tr class="a_b">
		//       <input name="a_b" type="checkbox" value="true"><span id="a_b">A B</span></input>
		//     </tr-->
		//   </body>
		// </html>`,
		// 			ExpectedJade: `html
		//   head
		//   body
		//     //
		//       tr class="a_b">
		//       <input name="a_b" type="checkbox" value="true"><span id="a_b">A B</span></input>
		//       </tr
		// `,
		// 			NilAssertion: assert.Nil,
		// 		},
		{
			Desc:    "TEST007 - Conditional",
			Skip:    doSKip,
			Options: defaultOptionsWithHead,
			SourceHTML: `<html>
  <head>
    <meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1" />
    <!--[if IE]>
      <script src="vendor/CFInstall.min.js"></script>
      <script src="javascripts/hello-msie.js"></script>
    <![endif]-->
    <meta http-equiv="content-type" content="text/html; charset=us-ascii" />
  </head>
  <body>
  <script>
//<![CDATA[
//]]>
  </script>
    Hello World.
    blah
    blah



    blah
  </body>
</html>`,
			ExpectedJade: `html
  head
    meta(http-equiv='X-UA-Compatible', content='IE=Edge,chrome=1')
    //if IE
      script(src='vendor/CFInstall.min.js')
      |       
      script(src='javascripts/hello-msie.js')
    |     
    meta(http-equiv='content-type', content='text/html; charset=us-ascii')
  |   
  body
    script.
      //<![CDATA[
      //]]>
    |     Hello World.
    |     blah
    |     blah
    |     blah
`,
			NilAssertion: assert.Nil,
		},
		{
			Desc:    "TEST008 - Conditional 2",
			Skip:    doSKip,
			Options: defaultOptionsWithHead,
			SourceHTML: `<!DOCTYPE html>

<!-- paulirish.com/2008/conditional-stylesheets-vs-css-hacks-answer-neither/ -->
<!--[if lt IE 7]> <html class="no-js lt-ie9 lt-ie8 lt-ie7" lang="en"> <![endif]-->
<!--[if IE 7]>    <html class="no-js lt-ie9 lt-ie8" lang="en"> <![endif]-->
<!--[if IE 8]>    <html class="no-js lt-ie9" lang="en"> <![endif]-->
<!--[if gt IE 8]><!--> <html class="no-js" lang="en"> <!--<![endif]-->
<head>
  <meta charset="utf-8" />

  <!-- Set the viewport width to device width for mobile -->
  <meta name="viewport" content="width=device-width" />

  <title>Welcome to Foundation</title>

  <!-- Included CSS Files -->
  <link rel="stylesheet" href="test/stylesheets/styles.css">

  <script src="vendor/assets/javascripts/foundation/modernizr.foundation.js"></script>

  <!-- IE Fix for HTML5 Tags -->
  <!--[if lt IE 9]>
    <script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
  <![endif]-->

</head>
<body>

  <div class="row">
    <div class="twelve columns">
      <h2>Welcome to Foundation</h2>
      <p>This is version 3.0.6 released on July 20, 2012.</p>
      <hr />
    </div>
  </div>

  <div class="row">
    <div class="eight columns">
      <h3>The Grid</h3>

      <!-- Grid Example -->
      <div class="row">
        <div class="twelve columns">
          <div class="panel">
            <p>This is a twelve column section in a row. Each of these includes a div.panel element so you can see where the columns are - it's not required at all for the grid.</p>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="six columns">
          <div class="panel">
            <p>Six columns</p>
          </div>
        </div>
        <div class="six columns">
          <div class="panel">
            <p>Six columns</p>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="four columns">
          <div class="panel">
            <p>Four columns</p>
          </div>
        </div>
        <div class="four columns">
          <div class="panel">
            <p>Four columns</p>
          </div>
        </div>
        <div class="four columns">
          <div class="panel">
            <p>Four columns</p>
          </div>
        </div>
      </div>

      <h3>Tabs</h3>
      <dl class="tabs">
        <dd class="active"><a href="#simple1">Simple Tab 1</a></dd>
        <dd><a href="#simple2">Simple Tab 2</a></dd>
        <dd><a href="#simple3">Simple Tab 3</a></dd>
      </dl>

      <ul class="tabs-content">
        <li class="active" id="simple1Tab">This is simple tab 1's content. Pretty neat, huh?</li>
        <li id="simple2Tab">This is simple tab 2's content. Now you see it!</li>
        <li id="simple3Tab">This is simple tab 3's content. It's, you know...okay.</li>
      </ul>

      <h3>Buttons</h3>

      <div class="row">
        <div class="six columns">
          <p><a href="#" class="small button">Small Button</a></p>
          <p><a href="#" class="button">Medium Button</a></p>
          <p><a href="#" class="large button">Large Button</a></p>
        </div>
        <div class="six columns">
          <p><a href="#" class="small alert button">Small Alert Button</a></p>
          <p><a href="#" class="success button">Medium Success Button</a></p>
          <p><a href="#" class="large secondary button">Large Secondary Button</a></p>
        </div>
      </div>
    </div>

    <div class="four columns">
      <h4>Getting Started</h4>
      <p>We're stoked you want to try Foundation! To get going, this file (index.html) includes some basic styles you can modify, play around with, or totally destroy to get going.</p>

      <h4>Other Resources</h4>
      <p>Once you've exhausted the fun in this document, you should check out:</p>
      <ul class="disc">
        <li><a href="http://foundation.zurb.com/docs">Foundation Documentation</a><br />Everything you need to know about using the framework.</li>
        <li><a href="http://github.com/zurb/foundation">Foundation on Github</a><br />Latest code, issue reports, feature requests and more.</li>
        <li><a href="http://twitter.com/foundationzurb">@foundationzurb</a><br />Ping us on Twitter if you have questions. If you build something with this we'd love to see it (and send you a totally boss sticker).</li>
      </ul>
    </div>
  </div>



  <!-- Included JS Files -->
  <script src="vendor/assets/javascripts/foundation/jquery.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.reveal.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.orbit.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.forms.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.placeholder.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.tooltips.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.alerts.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.buttons.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.accordion.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.navigation.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.mediaQueryToggle.js"></script>
  <script src="vendor/assets/javascripts/foundation/jquery.foundation.tabs.js"></script>
  <script src="vendor/assets/javascripts/foundation/app.js"></script>

</body>
</html>
`,
			ExpectedJade: `doctype html
// paulirish.com/2008/conditional-stylesheets-vs-css-hacks-answer-neither/
//if lt IE 7
  html.no-js.lt-ie9.lt-ie8.lt-ie7(lang='en')  
//if IE 7
  html.no-js.lt-ie9.lt-ie8(lang='en')  
//if IE 8
  html.no-js.lt-ie9(lang='en')  
// [if gt IE 8] <!
|  
html.no-js(lang='en')
  // <![endif]
  head
    meta(charset='utf-8')
    // Set the viewport width to device width for mobile
    meta(name='viewport', content='width=device-width')
    |   
    title Welcome to Foundation
    // Included CSS Files
    link(rel='stylesheet', href='test/stylesheets/styles.css')
    |   
    script(src='vendor/assets/javascripts/foundation/modernizr.foundation.js')
    // IE Fix for HTML5 Tags
    //if lt IE 9
      script(src='http://html5shiv.googlecode.com/svn/trunk/html5.js')
  body
    .row
      .twelve.columns
        h2 Welcome to Foundation
        |       
        p This is version 3.0.6 released on July 20, 2012.
        |       
        hr
    |   
    .row
      .eight.columns
        h3 The Grid
        // Grid Example
        .row
          .twelve.columns
            .panel
              p
                | This is a twelve column section in a row. Each of these includes a div.panel element so you can see where the columns are - it&apos;s not required at all for the grid.
        |       
        .row
          .six.columns
            .panel
              p Six columns
          |         
          .six.columns
            .panel
              p Six columns
        |       
        .row
          .four.columns
            .panel
              p Four columns
          |         
          .four.columns
            .panel
              p Four columns
          |         
          .four.columns
            .panel
              p Four columns
        |       
        h3 Tabs
        |       
        dl.tabs
          dd.active
            a(href='#simple1') Simple Tab 1
          |         
          dd
            a(href='#simple2') Simple Tab 2
          |         
          dd
            a(href='#simple3') Simple Tab 3
        |       
        ul.tabs-content
          li#simple1Tab.active This is simple tab 1&apos;s content. Pretty neat, huh?
          |         
          li#simple2Tab This is simple tab 2&apos;s content. Now you see it!
          |         
          li#simple3Tab This is simple tab 3&apos;s content. It&apos;s, you know...okay.
        |       
        h3 Buttons
        |       
        .row
          .six.columns
            p
              a.small.button(href='#') Small Button
            |           
            p
              a.button(href='#') Medium Button
            |           
            p
              a.large.button(href='#') Large Button
          |         
          .six.columns
            p
              a.small.alert.button(href='#') Small Alert Button
            |           
            p
              a.success.button(href='#') Medium Success Button
            |           
            p
              a.large.secondary.button(href='#') Large Secondary Button
      |     
      .four.columns
        h4 Getting Started
        |       
        p
          | We&apos;re stoked you want to try Foundation! To get going, this file (index.html) includes some basic styles you can modify, play around with, or totally destroy to get going.
        |       
        h4 Other Resources
        |       
        p Once you&apos;ve exhausted the fun in this document, you should check out:
        |       
        ul.disc
          li
            a(href='http://foundation.zurb.com/docs') Foundation Documentation
            br
            | Everything you need to know about using the framework.
          |         
          li
            a(href='http://github.com/zurb/foundation') Foundation on Github
            br
            | Latest code, issue reports, feature requests and more.
          |         
          li
            a(href='http://twitter.com/foundationzurb') @foundationzurb
            br
            | Ping us on Twitter if you have questions. If you build something with this we&apos;d love to see it (and send you a totally boss sticker).
    // Included JS Files
    script(src='vendor/assets/javascripts/foundation/jquery.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.reveal.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.orbit.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.forms.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.placeholder.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.tooltips.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.alerts.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.buttons.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.accordion.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.navigation.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.mediaQueryToggle.js')
    |   
    script(src='vendor/assets/javascripts/foundation/jquery.foundation.tabs.js')
    |   
    script(src='vendor/assets/javascripts/foundation/app.js')
`,
			NilAssertion: assert.Nil,
		},
		// {
		// 			Desc: "TEST009 - Conditional 3",
		// 			SourceHTML: `<!DOCTYPE html>
		// <!--[if lt IE 7]>      <html class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
		// <!--[if IE 7]>         <html class="no-js lt-ie9 lt-ie8"> <![endif]-->
		// <!--[if IE 8]>         <html class="no-js lt-ie9"> <![endif]-->
		// <!--[if gt IE 8]><!--> <html class="no-js"> <!--<![endif]-->
		//     <head>
		//         <meta charset="utf-8">
		//         <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
		//         <title></title>
		//         <meta name="description" content="">
		//         <meta name="viewport" content="width=device-width">

		//         <link rel="stylesheet" href="css/bootstrap.min.css">
		//         <style>
		//             body {
		//                 padding-top: 50px;
		//                 padding-bottom: 20px;
		//             }
		//         </style>
		//         <link rel="stylesheet" href="css/bootstrap-theme.min.css">
		//         <link rel="stylesheet" href="css/main.css">

		//         <script src="js/vendor/modernizr-2.6.2-respond-1.1.0.min.js"></script>
		//     </head>
		//     <body>
		//         <!--[if lt IE 7]>
		//             <p class="chromeframe">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> or <a href="http://www.google.com/chromeframe/?redirect=true">activate Google Chrome Frame</a> to improve your experience.</p>
		//         <![endif]-->
		//     <div class="navbar navbar-inverse navbar-fixed-top">
		// </div>
		// </body>
		// `,
		// 			ExpectedJade: `doctype html
		// //if lt IE 7
		//   html.no-js.lt-ie9.lt-ie8.lt-ie7
		// //if IE 7
		//   html.no-js.lt-ie9.lt-ie8
		// //if IE 8
		//   html.no-js.lt-ie9
		// // [if gt IE 8] <!
		// |
		// html.no-js
		//   // <![endif]
		//   head
		//     meta(charset='utf-8')
		//     |
		//     meta(http-equiv='X-UA-Compatible', content='IE=edge,chrome=1')
		//     |
		//     title
		//     |
		//     meta(name='description', content='')
		//     |
		//     meta(name='viewport', content='width=device-width')
		//     |
		//     link(rel='stylesheet', href='css/bootstrap.min.css')
		//     |
		//     style.
		//       body {
		//       padding-top: 50px;
		//       padding-bottom: 20px;
		//       }
		//     |
		//     link(rel='stylesheet', href='css/bootstrap-theme.min.css')
		//     |
		//     link(rel='stylesheet', href='css/main.css')
		//     |
		//     script(src='js/vendor/modernizr-2.6.2-respond-1.1.0.min.js')
		//   |
		//   body
		//     //if lt IE 7
		//       p.chromeframe
		//         | You are using an
		//         strong outdated
		//         |  browser. Please
		//         a(href='http://browsehappy.com/') upgrade your browser
		//         |  or
		//         a(href='http://www.google.com/chromeframe/?redirect=true') activate Google Chrome Frame
		//         |  to improve your experience.
		//     |
		//     .navbar.navbar-inverse.navbar-fixed-top
		// `,
		// 			NilAssertion: assert.Nil,
		// 		},
		{
			Desc:    "TEST010 - Empty class",
			Skip:    doSKip,
			Options: defaultOptionsWithHead,
			SourceHTML: `<html>
  <head>
  </head>
  <body class="">
    <code>
      blah
      blah
    </code>
  </body>
</html>
`,
			ExpectedJade: `html
  head
  |   
  body
    code
      | blah
      | blah
`,
			NilAssertion: assert.Nil,
		},
		{
			Desc:    "TEST011 - Entity",
			Skip:    doSKip,
			Options: defaultOptions,
			SourceHTML: `<p>Note the lack of the <code>&lt;meta name="viewport" content="width=device-width, initial-scale=1.0"&gt;</code>, which disables the zooming aspect of sites in mobile devices. In addition, we reset our container's width and are basically good to go.</p>
<p>&copy; Company 2013</p>
`,
			ExpectedJade: `html
  body
    p
      | Note the lack of the 
      code &lt;meta name=&quot;viewport&quot; content=&quot;width=device-width, initial-scale=1.0&quot;&gt;
      | , which disables the zooming aspect of sites in mobile devices. In addition, we reset our container&apos;s width and are basically good to go.
    p &copy; Company 2013
`,
			NilAssertion: assert.Nil,
		},
		{
			Desc:    "TEST012 - Headless",
			Skip:    doSKip,
			Options: defaultOptionsWithHead,
			SourceHTML: `<html>
  <head>
    <script type="text/javascript">window.location = "/newsite/";</script>
  </head>
</html>
`,
			ExpectedJade: `html
  head
    script(type='text/javascript').
      window.location = "/newsite/";
`,
			NilAssertion: assert.Nil,
		},
		{
			Desc:    "TEST013 - HTML5 Boilerplate",
			Skip:    doSKip,
			Options: defaultOptionsWithHead,
			SourceHTML: `<!doctype html>
<!--[if lt IE 7]> <html class="no-js ie6 oldie" lang="en"> <![endif]-->
<!--[if IE 7]>    <html class="no-js ie7 oldie" lang="en"> <![endif]-->
<!--[if IE 8]>    <html class="no-js ie8 oldie" lang="en"> <![endif]-->
<!--[if gt IE 8]><!--> <html class="no-js" lang="en"> <!--<![endif]-->
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">

	<title></title>
	<meta name="description" content="">
	<meta name="author" content="">

	<meta name="viewport" content="width=device-width,initial-scale=1">

	<link rel="stylesheet" href="css/style.css">

	<script src="js/libs/modernizr-2.0.6.min.js"></script>
</head>
<body>

<div id="container">
	<header>

	</header>
	<div id="main" role="main">

	</div>
	<footer>

	</footer>
</div> <!--! end of #container -->

<script src="//ajax.googleapis.com/ajax/libs/jquery/1.6.2/jquery.min.js"></script>
<script>window.jQuery || document.write('<script src="js/libs/jquery-1.6.2.min.js"><\/script>')</script>

<!-- scripts concatenated and minified via ant build script-->
<script src="js/plugins.js"></script>
<script src="js/script.js"></script>
<!-- end scripts-->

<script>
	var _gaq=[['_setAccount','UA-XXXXX-X'],['_trackPageview']]; // Change UA-XXXXX-X to be your site's ID
	(function(d,t){var g=d.createElement(t),s=d.getElementsByTagName(t)[0];g.async=1;
	g.src=('https:'==location.protocol?'//ssl':'//www')+'.google-analytics.com/ga.js';
	s.parentNode.insertBefore(g,s)}(document,'script'));
</script>

<!--[if lt IE 7 ]>
	<script src="//ajax.googleapis.com/ajax/libs/chrome-frame/1.0.2/CFInstall.min.js"></script>
	<script>window.attachEvent("onload",function(){CFInstall.check({mode:"overlay"})})</script>
<![endif]-->

</body>
</html>
`,
			ExpectedJade: `doctype html
//if lt IE 7
  html.no-js.ie6.oldie(lang='en')  
//if IE 7
  html.no-js.ie7.oldie(lang='en')  
//if IE 8
  html.no-js.ie8.oldie(lang='en')  
// [if gt IE 8] <!
|  
html.no-js(lang='en')
  // <![endif]
  head
    meta(charset='utf-8')
    | &#x9;
    meta(http-equiv='X-UA-Compatible', content='IE=edge,chrome=1')
    | &#x9;
    title
    | &#x9;
    meta(name='description', content='')
    | &#x9;
    meta(name='author', content='')
    | &#x9;
    meta(name='viewport', content='width=device-width,initial-scale=1')
    | &#x9;
    link(rel='stylesheet', href='css/style.css')
    | &#x9;
    script(src='js/libs/modernizr-2.0.6.min.js')
  body
    #container
      header
      | &#x9;
      #main(role='main')
      | &#x9;
      footer
    // ! end of #container
    script(src='//ajax.googleapis.com/ajax/libs/jquery/1.6.2/jquery.min.js')
    script.
      window.jQuery || document.write('<script src="js/libs/jquery-1.6.2.min.js"><\\/script>')
    // scripts concatenated and minified via ant build script
    script(src='js/plugins.js')
    script(src='js/script.js')
    // end scripts
    script.
      var _gaq=[['_setAccount','UA-XXXXX-X'],['_trackPageview']]; // Change UA-XXXXX-X to be your site's ID
      (function(d,t){var g=d.createElement(t),s=d.getElementsByTagName(t)[0];g.async=1;
      g.src=('https:'==location.protocol?'//ssl':'//www')+'.google-analytics.com/ga.js';
      s.parentNode.insertBefore(g,s)}(document,'script'));
    //if lt IE 7 
      script(src='//ajax.googleapis.com/ajax/libs/chrome-frame/1.0.2/CFInstall.min.js')
      | &#x9;
      script.
        window.attachEvent("onload",function(){CFInstall.check({mode:"overlay"})})
`,
			NilAssertion: assert.Nil,
		},
		{
			Desc:    "TEST014 - Leading Equal",
			Options: defaultOptions,
			SourceHTML: `<html><body><div>=1+1</div></body></html>
`,
			ExpectedJade: `html
  body
    div =1+1
`,
			NilAssertion: assert.Nil,
		},
		{
			Desc:    "TEST015 - Mustache",
			Options: defaultOptions,
			SourceHTML: `<div id="mustacheTestcases">
<div id="div1" class="panel-body {{listTypeClass}}">Test</div>
<div class="note checklist-part indent-{{indent}}"></div>
<textarea class="note-text" id="{{id}}">{{text}}</textarea>
</div>
`,
			ExpectedJade: `html
  body
    #mustacheTestcases
      #div1.panel-body(class='{{listTypeClass}}') Test
      .note.checklist-part(class='indent-{{indent}}')
      textarea.note-text(id='{{id}}') {{text}}
`,
			NilAssertion: assert.Nil,
		}, {
			Desc:    "TEST016 - Pfft Edge Case",
			Options: defaultOptions,
			SourceHTML: `<p>ffft</p>
`,
			ExpectedJade: `html
  body
    p ffft
`,
			NilAssertion: assert.Nil,
		},
		{
			Desc:    "TEST017 - Pre 1",
			Options: defaultOptionsWithHead,
			Skip:    doSKip,
			SourceHTML: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<html>
<head>
<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1">
<title>html2jade.js</title>
<link rel="stylesheet" type="text/css" href="highlight.css">
</head>
<body class="hl">
<pre class="hl"><span class="hl slc">// Generated by CoffeeScript 1.3.3</span>
<span class="hl opt">(</span><span class="hl kwa">function</span><span class="hl opt">() {</span>
  <span class="hl kwa">var</span> Converter<span class="hl opt">,</span> Output<span class="hl opt">,</span> Parser<span class="hl opt">,</span> StreamOutput<span class="hl opt">,</span> StringOutput<span class="hl opt">,</span> Writer<span class="hl opt">,</span> publicIdDocTypeNames<span class="hl opt">,</span> scope<span class="hl opt">,</span> systemIdDocTypeNames<span class="hl opt">,</span> _ref<span class="hl opt">,</span>
    __hasProp <span class="hl opt">= {}.</span>hasOwnProperty<span class="hl opt">,</span>
    __extends <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>child<span class="hl opt">,</span> parent<span class="hl opt">) {</span> <span class="hl kwa">for</span> <span class="hl opt">(</span><span class="hl kwa">var</span> key <span class="hl kwa">in</span> parent<span class="hl opt">) {</span> <span class="hl kwa">if</span> <span class="hl opt">(</span>__hasProp<span class="hl opt">.</span><span class="hl kwd">call</span><span class="hl opt">(</span>parent<span class="hl opt">,</span> key<span class="hl opt">))</span> child<span class="hl kwc">[key]</span> <span class="hl opt">=</span> parent<span class="hl kwc">[key]</span><span class="hl opt">; }</span> <span class="hl kwa">function</span> <span class="hl kwd">ctor</span><span class="hl opt">() {</span> <span class="hl kwa">this</span><span class="hl opt">.</span>constructor <span class="hl opt">=</span> child<span class="hl opt">; }</span> ctor<span class="hl opt">.</span><span class="hl kwa">prototype</span> <span class="hl opt">=</span> parent<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">;</span> child<span class="hl opt">.</span><span class="hl kwa">prototype</span> <span class="hl opt">=</span> <span class="hl kwa">new</span> <span class="hl kwd">ctor</span><span class="hl opt">();</span> child<span class="hl opt">.</span>__super__ <span class="hl opt">=</span> parent<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">;</span> <span class="hl kwa">return</span> child<span class="hl opt">; };</span>

  scope <span class="hl opt">=</span> <span class="hl kwa">typeof</span> exports <span class="hl opt">!==</span> <span class="hl str">&quot;undefined&quot;</span> <span class="hl opt">&amp;&amp;</span> exports <span class="hl opt">!==</span> <span class="hl kwa">null</span> ? exports <span class="hl opt">: (</span>_ref <span class="hl opt">=</span> <span class="hl kwa">this</span><span class="hl opt">.</span>Html2Jade<span class="hl opt">) !=</span> <span class="hl kwa">null</span> ? _ref <span class="hl opt">:</span> <span class="hl kwa">this</span><span class="hl opt">.</span>Html2Jade <span class="hl opt">= {};</span>

  Parser <span class="hl opt">= (</span><span class="hl kwa">function</span><span class="hl opt">() {</span>

    <span class="hl kwa">function</span> <span class="hl kwd">Parser</span><span class="hl opt">(</span>options<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>options <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        options <span class="hl opt">= {};</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">this</span><span class="hl opt">.</span>jsdom <span class="hl opt">=</span> <span class="hl kwd">require</span><span class="hl opt">(</span><span class="hl str">'jsdom'</span><span class="hl opt">);</span>
    <span class="hl opt">}</span>

    Parser<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>parse <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>arg<span class="hl opt">,</span> cb<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(!</span>arg<span class="hl opt">) {</span>
        <span class="hl kwa">return</span> <span class="hl kwd">cb</span><span class="hl opt">(</span><span class="hl str">'null file'</span><span class="hl opt">);</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>jsdom<span class="hl opt">.</span><span class="hl kwd">env</span><span class="hl opt">(</span>arg<span class="hl opt">,</span> cb<span class="hl opt">);</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    <span class="hl kwa">return</span> Parser<span class="hl opt">;</span>

  <span class="hl opt">})();</span>

  Writer <span class="hl opt">= (</span><span class="hl kwa">function</span><span class="hl opt">() {</span>

    <span class="hl kwa">function</span> <span class="hl kwd">Writer</span><span class="hl opt">(</span>options<span class="hl opt">) {</span>
      <span class="hl kwa">var</span> _ref1<span class="hl opt">,</span> _ref2<span class="hl opt">;</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>options <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        options <span class="hl opt">= {};</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">this</span><span class="hl opt">.</span>wrapLength <span class="hl opt">= (</span>_ref1 <span class="hl opt">=</span> options<span class="hl opt">.</span>wrapLength<span class="hl opt">) !=</span> <span class="hl kwa">null</span> ? _ref1 <span class="hl opt">:</span> <span class="hl num">80</span><span class="hl opt">;</span>
      <span class="hl kwa">this</span><span class="hl opt">.</span>scalate <span class="hl opt">= (</span>_ref2 <span class="hl opt">=</span> options<span class="hl opt">.</span>scalate<span class="hl opt">) !=</span> <span class="hl kwa">null</span> ? _ref2 <span class="hl opt">:</span> <span class="hl kwa">false</span><span class="hl opt">;</span>
      <span class="hl kwa">this</span><span class="hl opt">.</span>attrSep <span class="hl opt">=</span> <span class="hl kwa">this</span><span class="hl opt">.</span>scalate ? <span class="hl str">' '</span> <span class="hl opt">:</span> <span class="hl str">', '</span><span class="hl opt">;</span>
    <span class="hl opt">}</span>

    Writer<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>tagHead <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>node<span class="hl opt">) {</span>
      <span class="hl kwa">var</span> classes<span class="hl opt">,</span> result<span class="hl opt">;</span>
      result <span class="hl opt">=</span> node<span class="hl opt">.</span>tagName <span class="hl opt">!==</span> <span class="hl str">'DIV'</span> ? node<span class="hl opt">.</span>tagName<span class="hl opt">.</span><span class="hl kwd">toLowerCase</span><span class="hl opt">() :</span> <span class="hl str">''</span><span class="hl opt">;</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>node<span class="hl opt">.</span>id<span class="hl opt">) {</span>
        result <span class="hl opt">+=</span> <span class="hl str">'#'</span> <span class="hl opt">+</span> node<span class="hl opt">.</span>id<span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>node<span class="hl opt">.</span><span class="hl kwd">hasAttribute</span><span class="hl opt">(</span><span class="hl str">'class'</span><span class="hl opt">) &amp;&amp;</span> node<span class="hl opt">.</span><span class="hl kwd">getAttribute</span><span class="hl opt">(</span><span class="hl str">'class'</span><span class="hl opt">).</span>length <span class="hl opt">&gt;</span> <span class="hl num">0</span><span class="hl opt">) {</span>
        classes <span class="hl opt">=</span> node<span class="hl opt">.</span><span class="hl kwd">getAttribute</span><span class="hl opt">(</span><span class="hl str">'class'</span><span class="hl opt">).</span><span class="hl kwd">split</span><span class="hl opt">(/</span>\s<span class="hl opt">+/).</span><span class="hl kwd">filter</span><span class="hl opt">(</span><span class="hl kwa">function</span><span class="hl opt">(</span>item<span class="hl opt">) {</span>
          <span class="hl kwa">return</span> <span class="hl opt">(</span>item <span class="hl opt">!=</span> <span class="hl kwa">null</span><span class="hl opt">) &amp;&amp;</span> item<span class="hl opt">.</span><span class="hl kwd">trim</span><span class="hl opt">().</span>length <span class="hl opt">&gt;</span> <span class="hl num">0</span><span class="hl opt">;</span>
        <span class="hl opt">});</span>
        result <span class="hl opt">+=</span> <span class="hl str">'.'</span> <span class="hl opt">+</span> classes<span class="hl opt">.</span><span class="hl kwd">join</span><span class="hl opt">(</span><span class="hl str">'.'</span><span class="hl opt">);</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>result<span class="hl opt">.</span>length <span class="hl opt">===</span> <span class="hl num">0</span><span class="hl opt">) {</span>
        result <span class="hl opt">=</span> <span class="hl str">'div'</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">return</span> result<span class="hl opt">;</span>
    <span class="hl opt">};</span>

    Writer<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>tagAttr <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>node<span class="hl opt">) {</span>
      <span class="hl kwa">var</span> attr<span class="hl opt">,</span> attrs<span class="hl opt">,</span> nodeName<span class="hl opt">,</span> result<span class="hl opt">,</span> _i<span class="hl opt">,</span> _len<span class="hl opt">;</span>
      attrs <span class="hl opt">=</span> node<span class="hl opt">.</span>attributes<span class="hl opt">;</span>
      <span class="hl kwa">if</span> <span class="hl opt">(!</span>attrs || attrs<span class="hl opt">.</span>length <span class="hl opt">===</span> <span class="hl num">0</span><span class="hl opt">) {</span>
        <span class="hl kwa">return</span> <span class="hl str">''</span><span class="hl opt">;</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        result <span class="hl opt">= [];</span>
        <span class="hl kwa">for</span> <span class="hl opt">(</span>_i <span class="hl opt">=</span> <span class="hl num">0</span><span class="hl opt">,</span> _len <span class="hl opt">=</span> attrs<span class="hl opt">.</span>length<span class="hl opt">;</span> _i <span class="hl opt">&lt;</span> _len<span class="hl opt">;</span> _i<span class="hl opt">++) {</span>
          attr <span class="hl opt">=</span> attrs<span class="hl kwc">[_i]</span><span class="hl opt">;</span>
          <span class="hl kwa">if</span> <span class="hl opt">(</span>attr <span class="hl opt">&amp;&amp; (</span>nodeName <span class="hl opt">=</span> attr<span class="hl opt">.</span>nodeName<span class="hl opt">)) {</span>
            <span class="hl kwa">if</span> <span class="hl opt">(</span>nodeName <span class="hl opt">!==</span> <span class="hl str">'id'</span> <span class="hl opt">&amp;&amp;</span> nodeName <span class="hl opt">!==</span> <span class="hl str">'class'</span> <span class="hl opt">&amp;&amp;</span> <span class="hl kwa">typeof</span> <span class="hl opt">(</span>attr<span class="hl opt">.</span>nodeValue <span class="hl opt">!=</span> <span class="hl kwa">null</span><span class="hl opt">)) {</span>
              result<span class="hl opt">.</span><span class="hl kwd">push</span><span class="hl opt">(</span>attr<span class="hl opt">.</span>nodeName <span class="hl opt">+</span> <span class="hl str">'=</span><span class="hl esc">\'</span><span class="hl str">'</span> <span class="hl opt">+</span> attr<span class="hl opt">.</span>nodeValue<span class="hl opt">.</span><span class="hl kwd">replace</span><span class="hl opt">(/</span><span class="hl str">'/g, '</span><span class="hl esc">\\\'</span><span class="hl str">') + '</span><span class="hl esc">\'</span><span class="hl str">');</span>
<span class="hl str">            }</span>
<span class="hl str">          }</span>
<span class="hl str">        }</span>
<span class="hl str">        if (result.length &gt; 0) {</span>
<span class="hl str">          return '</span><span class="hl opt">(</span><span class="hl str">' + result.join(this.attrSep) + '</span><span class="hl opt">)</span><span class="hl str">';</span>
<span class="hl str">        } else {</span>
<span class="hl str">          return '</span><span class="hl str">';</span>
<span class="hl str">        }</span>
<span class="hl str">      }</span>
<span class="hl str">    };</span>
<span class="hl str"></span>
<span class="hl str">    Writer.prototype.tagText = function(node) {</span>
<span class="hl str">      var data, _ref1;</span>
<span class="hl str">      if (((_ref1 = node.firstChild) != null ? _ref1.nodeType : void 0) !== 3) {</span>
<span class="hl str">        return null;</span>
<span class="hl str">      } else if (node.firstChild !== node.lastChild) {</span>
<span class="hl str">        return null;</span>
<span class="hl str">      } else {</span>
<span class="hl str">        data = node.firstChild.data;</span>
<span class="hl str">        if (data.length &gt; this.wrapLength || data.match(/</span><span class="hl esc">\r</span><span class="hl str">|</span><span class="hl esc">\n</span><span class="hl str">/)) {</span>
<span class="hl str">          return null;</span>
<span class="hl str">        } else {</span>
<span class="hl str">          return data;</span>
<span class="hl str">        }</span>
<span class="hl str">      }</span>
<span class="hl str">    };</span>
<span class="hl str"></span>
<span class="hl str">    Writer.prototype.forEachChild = function(parent, cb) {</span>
<span class="hl str">      var child, _results;</span>
<span class="hl str">      if (parent) {</span>
<span class="hl str">        child = parent.firstChild;</span>
<span class="hl str">        _results = [];</span>
<span class="hl str">        while (child) {</span>
<span class="hl str">          cb(child);</span>
<span class="hl str">          _results.push(child = child.nextSibling);</span>
<span class="hl str">        }</span>
<span class="hl str">        return _results;</span>
<span class="hl str">      }</span>
<span class="hl str">    };</span>
<span class="hl str"></span>
<span class="hl str">    Writer.prototype.writeTextContent = function(node, output, pipe, trim, wrap, escapeBackslash) {</span>
<span class="hl str">      var _this = this;</span>
<span class="hl str">      if (pipe == null) {</span>
<span class="hl str">        pipe = true;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (trim == null) {</span>
<span class="hl str">        trim = true;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (wrap == null) {</span>
<span class="hl str">        wrap = true;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (escapeBackslash == null) {</span>
<span class="hl str">        escapeBackslash = false;</span>
<span class="hl str">      }</span>
<span class="hl str">      output.enter();</span>
<span class="hl str">      this.forEachChild(node, function(child) {</span>
<span class="hl str">        return _this.writeText(child, output, pipe, trim, wrap, escapeBackslash);</span>
<span class="hl str">      });</span>
<span class="hl str">      return output.leave();</span>
<span class="hl str">    };</span>
<span class="hl str"></span>
<span class="hl str">    Writer.prototype.writeText = function(node, output, pipe, trim, wrap, escapeBackslash) {</span>
<span class="hl str">      var data, lines,</span>
<span class="hl str">        _this = this;</span>
<span class="hl str">      if (pipe == null) {</span>
<span class="hl str">        pipe = true;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (trim == null) {</span>
<span class="hl str">        trim = true;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (wrap == null) {</span>
<span class="hl str">        wrap = true;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (escapeBackslash == null) {</span>
<span class="hl str">        escapeBackslash = false;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (node.nodeType === 3) {</span>
<span class="hl str">        data = node.data || '</span><span class="hl str">';</span>
<span class="hl str">        if (data.length &gt; 0) {</span>
<span class="hl str">          lines = data.split(/</span><span class="hl esc">\r</span><span class="hl str">|</span><span class="hl esc">\n</span><span class="hl str">/);</span>
<span class="hl str">          return lines.forEach(function(line) {</span>
<span class="hl str">            return _this.writeTextLine(line, output, pipe, trim, wrap, escapeBackslash);</span>
<span class="hl str">          });</span>
<span class="hl str">        }</span>
<span class="hl str">      }</span>
<span class="hl str">    };</span>
<span class="hl str"></span>
<span class="hl str">    Writer.prototype.writeTextLine = function(line, output, pipe, trim, wrap, escapeBackslash) {</span>
<span class="hl str">      var lines, prefix,</span>
<span class="hl str">        _this = this;</span>
<span class="hl str">      if (pipe == null) {</span>
<span class="hl str">        pipe = true;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (trim == null) {</span>
<span class="hl str">        trim = true;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (wrap == null) {</span>
<span class="hl str">        wrap = true;</span>
<span class="hl str">      }</span>
<span class="hl str">      if (escapeBackslash == null) {</span>
<span class="hl str">        escapeBackslash = false;</span>
<span class="hl str">      }</span>
<span class="hl str">      prefix = pipe ? '</span>| <span class="hl str">' : '</span><span class="hl str">';</span>
<span class="hl str">      if (trim) {</span>
<span class="hl str">        line = line ? line.trim() : '</span><span class="hl str">';</span>
<span class="hl str">      }</span>
<span class="hl str">      if (line &amp;&amp; line.length &gt; 0) {</span>
<span class="hl str">        if (escapeBackslash) {</span>
<span class="hl str">          line = line.replace(&quot;</span><span class="hl esc">\\</span><span class="hl str">&quot;, &quot;</span><span class="hl esc">\\\\</span><span class="hl str">&quot;);</span>
<span class="hl str">        }</span>
<span class="hl str">        if (!wrap || line.length &lt;= this.wrapLength) {</span>
<span class="hl str">          return output.writeln(prefix + line);</span>
<span class="hl str">        } else {</span>
<span class="hl str">          lines = this.breakLine(line);</span>
<span class="hl str">          if (lines.length === 1) {</span>
<span class="hl str">            return output.writeln(prefix + line);</span>
<span class="hl str">          } else {</span>
<span class="hl str">            return lines.forEach(function(line) {</span>
<span class="hl str">              return _this.writeTextLine(line, output, pipe, trim, wrap);</span>
<span class="hl str">            });</span>
<span class="hl str">          }</span>
<span class="hl str">        }</span>
<span class="hl str">      }</span>
<span class="hl str">    };</span>
<span class="hl str"></span>
<span class="hl str">    Writer.prototype.breakLine = function(line) {</span>
<span class="hl str">      var lines, word, words;</span>
<span class="hl str">      if (!line || line.length === 0) {</span>
<span class="hl str">        return [];</span>
<span class="hl str">      }</span>
<span class="hl str">      if (line.search(/\s+/ === -1)) {</span>
<span class="hl str">        return [line];</span>
<span class="hl str">      }</span>
<span class="hl str">      lines = [];</span>
<span class="hl str">      words = line.split(/\s+/);</span>
<span class="hl str">      line = '</span><span class="hl str">';</span>
<span class="hl str">      while (words.length) {</span>
<span class="hl str">        word = words.shift();</span>
<span class="hl str">        if (line.length + word.length &gt; this.wrapLength) {</span>
<span class="hl str">          lines.push(line);</span>
<span class="hl str">          line = word;</span>
<span class="hl str">        } else if (line.length) {</span>
<span class="hl str">          line += '</span> <span class="hl str">' + word;</span>
<span class="hl str">        } else {</span>
<span class="hl str">          line = word;</span>
<span class="hl str">        }</span>
<span class="hl str">      }</span>
<span class="hl str">      if (line.length) {</span>
<span class="hl str">        lines.push(line);</span>
<span class="hl str">      }</span>
<span class="hl str">      return lines;</span>
<span class="hl str">    };</span>
<span class="hl str"></span>
<span class="hl str">    return Writer;</span>
<span class="hl str"></span>
<span class="hl str">  })();</span>
<span class="hl str"></span>
<span class="hl str">  publicIdDocTypeNames = {</span>
<span class="hl str">    &quot;-//W3C//DTD XHTML 1.0 Transitional//EN&quot;: &quot;transitional&quot;,</span>
<span class="hl str">    &quot;-//W3C//DTD XHTML 1.0 Strict//EN&quot;: &quot;strict&quot;,</span>
<span class="hl str">    &quot;-//W3C//DTD XHTML 1.0 Frameset//EN&quot;: &quot;frameset&quot;,</span>
<span class="hl str">    &quot;-//W3C//DTD XHTML 1.1//EN&quot;: &quot;1.1&quot;,</span>
<span class="hl str">    &quot;-//W3C//DTD XHTML Basic 1.1//EN&quot;: &quot;basic&quot;,</span>
<span class="hl str">    &quot;-//WAPFORUM//DTD XHTML Mobile 1.2//EN&quot;: &quot;mobile&quot;</span>
<span class="hl str">  };</span>
<span class="hl str"></span>
<span class="hl str">  systemIdDocTypeNames = {</span>
<span class="hl str">    &quot;http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd&quot;: &quot;transitional&quot;,</span>
<span class="hl str">    &quot;http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd&quot;: &quot;strict&quot;,</span>
<span class="hl str">    &quot;http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd&quot;: &quot;frameset&quot;,</span>
<span class="hl str">    &quot;http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd&quot;: &quot;1.1&quot;,</span>
<span class="hl str">    &quot;http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd&quot;: &quot;basic&quot;,</span>
<span class="hl str">    &quot;http://www.openmobilealliance.org/tech/DTD/xhtml-mobile12.dtd&quot;: &quot;mobile&quot;</span>
<span class="hl str">  };</span>
<span class="hl str"></span>
<span class="hl str">  Converter = (function() {</span>
<span class="hl str"></span>
<span class="hl str">    function Converter(options) {</span>
<span class="hl str">      var _ref1, _ref2;</span>
<span class="hl str">      if (options == null) {</span>
<span class="hl str">        options = {};</span>
<span class="hl str">      }</span>
<span class="hl str">      this.scalate = (_ref1 = options.scalate) != null ? _ref1 : false;</span>
<span class="hl str">      this.writer = (_ref2 = options.writer) != null ? _ref2 : new Writer(options);</span>
<span class="hl str">    }</span>
<span class="hl str"></span>
<span class="hl str">    Converter.prototype.document = function(document, output) {</span>
<span class="hl str">      var docTypeName, doctype, htmlEls, publicId, systemId;</span>
<span class="hl str">      if (document.doctype != null) {</span>
<span class="hl str">        doctype = document.doctype;</span>
<span class="hl str">        docTypeName = void 0;</span>
<span class="hl str">        publicId = doctype.publicId;</span>
<span class="hl str">        systemId = doctype.systemId;</span>
<span class="hl str">        if ((publicId != null) &amp;&amp; (publicIdDocTypeNames[publicId] != null)) {</span>
<span class="hl str">          docTypeName = publicIdDocTypeNames[publicId];</span>
<span class="hl str">        } else if ((systemId != null) &amp;&amp; (systemIdDocTypeNames[systemId] != null)) {</span>
<span class="hl str">          docTypeName = systemIdDocTypeNames[systemId] != null;</span>
<span class="hl str">        } else if ((doctype.name != null) &amp;&amp; doctype.name.toLowerCase() === '</span>html<span class="hl str">') {</span>
<span class="hl str">          docTypeName = '</span><span class="hl num">5</span><span class="hl str">';</span>
<span class="hl str">        }</span>
<span class="hl str">        if (docTypeName != null) {</span>
<span class="hl str">          output.writeln('</span><span class="hl opt">!!!</span> <span class="hl str">' + docTypeName);</span>
<span class="hl str">        }</span>
<span class="hl str">      }</span>
<span class="hl str">      if (document.documentElement) {</span>
<span class="hl str">        return this.children(document, output, false);</span>
<span class="hl str">      } else {</span>
<span class="hl str">        htmlEls = document.getElementsByTagName('</span>html<span class="hl str">');</span>
<span class="hl str">        if (htmlEls.length &gt; 0) {</span>
<span class="hl str">          return this.element(htmlEls[0], output);</span>
<span class="hl str">        }</span>
<span class="hl str">      }</span>
<span class="hl str">    };</span>
<span class="hl str"></span>
<span class="hl str">    Converter.prototype.element = function(node, output) {</span>
<span class="hl str">      var firstline, tagAttr, tagHead, tagName, tagText,</span>
<span class="hl str">        _this = this;</span>
<span class="hl str">      if (!(node != null ? node.tagName : void 0)) {</span>
<span class="hl str">        return;</span>
<span class="hl str">      }</span>
<span class="hl str">      tagName = node.tagName.toLowerCase();</span>
<span class="hl str">      tagHead = this.writer.tagHead(node);</span>
<span class="hl str">      tagAttr = this.writer.tagAttr(node);</span>
<span class="hl str">      tagText = this.writer.tagText(node);</span>
<span class="hl str">      if (tagName === '</span>script<span class="hl str">' || tagName === '</span>style<span class="hl str">') {</span>
<span class="hl str">        if (node.hasAttribute('</span>src<span class="hl str">')) {</span>
<span class="hl str">          output.writeln(tagHead + tagAttr);</span>
<span class="hl str">          return this.writer.writeTextContent(node, output, false, false, false);</span>
<span class="hl str">        } else if (tagName === '</span>script<span class="hl str">') {</span>
<span class="hl str">          return this.script(node, output, tagHead, tagAttr);</span>
<span class="hl str">        } else if (tagName === '</span>style<span class="hl str">') {</span>
<span class="hl str">          return this.style(node, output, tagHead, tagAttr);</span>
<span class="hl str">        }</span>
<span class="hl str">      } else if (tagName === '</span>conditional<span class="hl str">') {</span>
<span class="hl str">        output.writeln('</span><span class="hl slc">//' + node.getAttribute('condition'));</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span><span class="hl kwd">children</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">);</span>
      <span class="hl opt">}</span> <span class="hl kwa">else if</span> <span class="hl opt">([</span><span class="hl str">'pre'</span><span class="hl opt">].</span><span class="hl kwd">indexOf</span><span class="hl opt">(</span>tagName<span class="hl opt">) !== -</span><span class="hl num">1</span><span class="hl opt">) {</span>
        output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">(</span>tagHead <span class="hl opt">+</span> tagAttr <span class="hl opt">+</span> <span class="hl str">'.'</span><span class="hl opt">);</span>
        output<span class="hl opt">.</span><span class="hl kwd">enter</span><span class="hl opt">();</span>
        firstline <span class="hl opt">=</span> <span class="hl kwa">true</span><span class="hl opt">;</span>
        <span class="hl kwa">this</span><span class="hl opt">.</span>writer<span class="hl opt">.</span><span class="hl kwd">forEachChild</span><span class="hl opt">(</span>node<span class="hl opt">,</span> <span class="hl kwa">function</span><span class="hl opt">(</span>child<span class="hl opt">) {</span>
          <span class="hl kwa">var</span> data<span class="hl opt">;</span>
          <span class="hl kwa">if</span> <span class="hl opt">(</span>child<span class="hl opt">.</span>nodeType <span class="hl opt">===</span> <span class="hl num">3</span><span class="hl opt">) {</span>
            data <span class="hl opt">=</span> child<span class="hl opt">.</span>data<span class="hl opt">;</span>
            <span class="hl kwa">if</span> <span class="hl opt">((</span>data <span class="hl opt">!=</span> <span class="hl kwa">null</span><span class="hl opt">) &amp;&amp;</span> data<span class="hl opt">.</span>length <span class="hl opt">&gt;</span> <span class="hl num">0</span><span class="hl opt">) {</span>
              <span class="hl kwa">if</span> <span class="hl opt">(</span>firstline<span class="hl opt">) {</span>
                <span class="hl kwa">if</span> <span class="hl opt">(</span>data<span class="hl opt">.</span><span class="hl kwd">search</span><span class="hl opt">(/</span><span class="hl esc">\r\n</span>|<span class="hl esc">\r</span>|<span class="hl esc">\n</span><span class="hl opt">/) ===</span> <span class="hl num">0</span><span class="hl opt">) {</span>
                  data <span class="hl opt">=</span> data<span class="hl opt">.</span><span class="hl kwd">replace</span><span class="hl opt">(/</span><span class="hl esc">\r\n</span>|<span class="hl esc">\r</span>|<span class="hl esc">\n</span><span class="hl opt">/,</span> <span class="hl str">''</span><span class="hl opt">);</span>
                <span class="hl opt">}</span>
                data <span class="hl opt">=</span> <span class="hl str">'</span><span class="hl esc">\\</span><span class="hl str">n'</span> <span class="hl opt">+</span> data<span class="hl opt">;</span>
                firstline <span class="hl opt">=</span> <span class="hl kwa">false</span><span class="hl opt">;</span>
              <span class="hl opt">}</span>
              data <span class="hl opt">=</span> data<span class="hl opt">.</span><span class="hl kwd">replace</span><span class="hl opt">(/</span><span class="hl esc">\t</span><span class="hl opt">/</span>g<span class="hl opt">,</span> <span class="hl str">'</span><span class="hl esc">\\</span><span class="hl str">t'</span><span class="hl opt">);</span>
              data <span class="hl opt">=</span> data<span class="hl opt">.</span><span class="hl kwd">replace</span><span class="hl opt">(/</span><span class="hl esc">\r\n</span>|<span class="hl esc">\r</span>|<span class="hl esc">\n</span><span class="hl opt">/</span>g<span class="hl opt">,</span> <span class="hl str">'</span><span class="hl esc">\n</span><span class="hl str">'</span> <span class="hl opt">+</span> output<span class="hl opt">.</span>indents<span class="hl opt">);</span>
              <span class="hl kwa">return</span> output<span class="hl opt">.</span><span class="hl kwd">write</span><span class="hl opt">(</span>data<span class="hl opt">);</span>
            <span class="hl opt">}</span>
          <span class="hl opt">}</span>
        <span class="hl opt">});</span>
        output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">();</span>
        <span class="hl kwa">return</span> output<span class="hl opt">.</span><span class="hl kwd">leave</span><span class="hl opt">();</span>
      <span class="hl opt">}</span> <span class="hl kwa">else if</span> <span class="hl opt">(</span>tagText<span class="hl opt">) {</span>
        <span class="hl kwa">if</span> <span class="hl opt">(</span>tagText<span class="hl opt">.</span>length <span class="hl opt">&gt;</span> <span class="hl num">0</span> <span class="hl opt">&amp;&amp;</span> tagText<span class="hl opt">.</span><span class="hl kwd">charAt</span><span class="hl opt">(</span><span class="hl num">0</span><span class="hl opt">) ===</span> <span class="hl str">'='</span><span class="hl opt">) {</span>
          tagText <span class="hl opt">=</span> <span class="hl str">'</span><span class="hl esc">\\</span><span class="hl str">'</span> <span class="hl opt">+</span> tagText<span class="hl opt">;</span>
        <span class="hl opt">}</span>
        <span class="hl kwa">return</span> output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">(</span>tagHead <span class="hl opt">+</span> tagAttr <span class="hl opt">+</span> <span class="hl str">' '</span> <span class="hl opt">+</span> tagText<span class="hl opt">);</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">(</span>tagHead <span class="hl opt">+</span> tagAttr<span class="hl opt">);</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span><span class="hl kwd">children</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">);</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    Converter<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>children <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>parent<span class="hl opt">,</span> output<span class="hl opt">,</span> indent<span class="hl opt">) {</span>
      <span class="hl kwa">var</span> _this <span class="hl opt">=</span> <span class="hl kwa">this</span><span class="hl opt">;</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        indent <span class="hl opt">=</span> <span class="hl kwa">true</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent<span class="hl opt">) {</span>
        output<span class="hl opt">.</span><span class="hl kwd">enter</span><span class="hl opt">();</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">this</span><span class="hl opt">.</span>writer<span class="hl opt">.</span><span class="hl kwd">forEachChild</span><span class="hl opt">(</span>parent<span class="hl opt">,</span> <span class="hl kwa">function</span><span class="hl opt">(</span>child<span class="hl opt">) {</span>
        <span class="hl kwa">var</span> nodeType<span class="hl opt">;</span>
        nodeType <span class="hl opt">=</span> child<span class="hl opt">.</span>nodeType<span class="hl opt">;</span>
        <span class="hl kwa">if</span> <span class="hl opt">(</span>nodeType <span class="hl opt">===</span> <span class="hl num">1</span><span class="hl opt">) {</span>
          <span class="hl kwa">return</span> _this<span class="hl opt">.</span><span class="hl kwd">element</span><span class="hl opt">(</span>child<span class="hl opt">,</span> output<span class="hl opt">);</span>
        <span class="hl opt">}</span> <span class="hl kwa">else if</span> <span class="hl opt">(</span>nodeType <span class="hl opt">===</span> <span class="hl num">3</span><span class="hl opt">) {</span>
          <span class="hl kwa">if</span> <span class="hl opt">(</span>parent<span class="hl opt">.</span>_nodeName <span class="hl opt">===</span> <span class="hl str">'code'</span><span class="hl opt">) {</span>
            <span class="hl kwa">return</span> _this<span class="hl opt">.</span><span class="hl kwd">text</span><span class="hl opt">(</span>child<span class="hl opt">,</span> output<span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">true</span><span class="hl opt">,</span> <span class="hl kwa">true</span><span class="hl opt">);</span>
          <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
            <span class="hl kwa">return</span> _this<span class="hl opt">.</span><span class="hl kwd">text</span><span class="hl opt">(</span>child<span class="hl opt">,</span> output<span class="hl opt">);</span>
          <span class="hl opt">}</span>
        <span class="hl opt">}</span> <span class="hl kwa">else if</span> <span class="hl opt">(</span>nodeType <span class="hl opt">===</span> <span class="hl num">8</span><span class="hl opt">) {</span>
          <span class="hl kwa">return</span> _this<span class="hl opt">.</span><span class="hl kwd">comment</span><span class="hl opt">(</span>child<span class="hl opt">,</span> output<span class="hl opt">);</span>
        <span class="hl opt">}</span>
      <span class="hl opt">});</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent<span class="hl opt">) {</span>
        <span class="hl kwa">return</span> output<span class="hl opt">.</span><span class="hl kwd">leave</span><span class="hl opt">();</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    Converter<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>text <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">,</span> pipe<span class="hl opt">,</span> trim<span class="hl opt">,</span> wrap<span class="hl opt">) {</span>
      node<span class="hl opt">.</span><span class="hl kwd">normalize</span><span class="hl opt">();</span>
      <span class="hl kwa">return this</span><span class="hl opt">.</span>writer<span class="hl opt">.</span><span class="hl kwd">writeText</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">,</span> pipe<span class="hl opt">,</span> trim<span class="hl opt">,</span> wrap<span class="hl opt">);</span>
    <span class="hl opt">};</span>

    Converter<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>comment <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">) {</span>
      <span class="hl kwa">var</span> condition<span class="hl opt">,</span> data<span class="hl opt">,</span> lines<span class="hl opt">,</span>
        _this <span class="hl opt">=</span> <span class="hl kwa">this</span><span class="hl opt">;</span>
      condition <span class="hl opt">=</span> node<span class="hl opt">.</span>data<span class="hl opt">.</span><span class="hl kwd">match</span><span class="hl opt">(/</span>\s<span class="hl opt">*</span>\<span class="hl opt">[(</span><span class="hl kwa">if</span>\s<span class="hl opt">+[</span>^\<span class="hl opt">]]+)</span>\<span class="hl opt">]/);</span>
      <span class="hl kwa">if</span> <span class="hl opt">(!</span>condition<span class="hl opt">) {</span>
        data <span class="hl opt">=</span> node<span class="hl opt">.</span>data || <span class="hl str">''</span><span class="hl opt">;</span>
        <span class="hl kwa">if</span> <span class="hl opt">(</span>data<span class="hl opt">.</span>length <span class="hl opt">===</span> <span class="hl num">0</span> || data<span class="hl opt">.</span><span class="hl kwd">search</span><span class="hl opt">(/</span><span class="hl esc">\r</span>|<span class="hl esc">\n</span><span class="hl opt">/) === -</span><span class="hl num">1</span><span class="hl opt">) {</span>
          <span class="hl kwa">return</span> output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">(</span><span class="hl str">&quot;// &quot;</span> <span class="hl opt">+ (</span>data<span class="hl opt">.</span><span class="hl kwd">trim</span><span class="hl opt">()));</span>
        <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
          output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">(</span><span class="hl str">'//'</span><span class="hl opt">);</span>
          output<span class="hl opt">.</span><span class="hl kwd">enter</span><span class="hl opt">();</span>
          lines <span class="hl opt">=</span> data<span class="hl opt">.</span><span class="hl kwd">split</span><span class="hl opt">(/</span><span class="hl esc">\r</span>|<span class="hl esc">\n</span><span class="hl opt">/);</span>
          lines<span class="hl opt">.</span><span class="hl kwd">forEach</span><span class="hl opt">(</span><span class="hl kwa">function</span><span class="hl opt">(</span>line<span class="hl opt">) {</span>
            <span class="hl kwa">return</span> _this<span class="hl opt">.</span>writer<span class="hl opt">.</span><span class="hl kwd">writeTextLine</span><span class="hl opt">(</span>line<span class="hl opt">,</span> output<span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">);</span>
          <span class="hl opt">});</span>
          <span class="hl kwa">return</span> output<span class="hl opt">.</span><span class="hl kwd">leave</span><span class="hl opt">();</span>
        <span class="hl opt">}</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span><span class="hl kwd">conditional</span><span class="hl opt">(</span>node<span class="hl opt">,</span> condition<span class="hl kwc">[1]</span><span class="hl opt">,</span> output<span class="hl opt">);</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    Converter<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>conditional <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>node<span class="hl opt">,</span> condition<span class="hl opt">,</span> output<span class="hl opt">) {</span>
      <span class="hl kwa">var</span> conditionalElem<span class="hl opt">,</span> innerHTML<span class="hl opt">;</span>
      innerHTML <span class="hl opt">=</span> node<span class="hl opt">.</span>textContent<span class="hl opt">.</span><span class="hl kwd">trim</span><span class="hl opt">().</span><span class="hl kwd">replace</span><span class="hl opt">(/</span>\s<span class="hl opt">*</span>\<span class="hl opt">[</span><span class="hl kwa">if</span>\s<span class="hl opt">+[</span>^\<span class="hl opt">]]+</span>\<span class="hl opt">]&gt;</span>\s*/<span class="hl opt">,</span> <span class="hl str">''</span><span class="hl opt">).</span><span class="hl kwd">replace</span><span class="hl opt">(</span><span class="hl str">'&lt;![endif]'</span><span class="hl opt">,</span> <span class="hl str">''</span><span class="hl opt">);</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>innerHTML<span class="hl opt">.</span><span class="hl kwd">indexOf</span><span class="hl opt">(</span><span class="hl str">&quot;&lt;!&quot;</span><span class="hl opt">) ===</span> <span class="hl num">0</span><span class="hl opt">) {</span>
        condition <span class="hl opt">=</span> <span class="hl str">&quot; [&quot;</span> <span class="hl opt">+</span> condition <span class="hl opt">+</span> <span class="hl str">&quot;] &lt;!&quot;</span><span class="hl opt">;</span>
        innerHTML <span class="hl opt">=</span> <span class="hl kwa">null</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      conditionalElem <span class="hl opt">=</span> node<span class="hl opt">.</span>ownerDocument<span class="hl opt">.</span><span class="hl kwd">createElement</span><span class="hl opt">(</span><span class="hl str">'conditional'</span><span class="hl opt">);</span>
      conditionalElem<span class="hl opt">.</span><span class="hl kwd">setAttribute</span><span class="hl opt">(</span><span class="hl str">'condition'</span><span class="hl opt">,</span> condition<span class="hl opt">);</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>innerHTML<span class="hl opt">) {</span>
        conditionalElem<span class="hl opt">.</span>innerHTML <span class="hl opt">=</span> innerHTML<span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">return</span> node<span class="hl opt">.</span>parentNode<span class="hl opt">.</span><span class="hl kwd">insertBefore</span><span class="hl opt">(</span>conditionalElem<span class="hl opt">,</span> node<span class="hl opt">.</span>nextSibling<span class="hl opt">);</span>
    <span class="hl opt">};</span>

    Converter<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>script <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">,</span> tagHead<span class="hl opt">,</span> tagAttr<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span><span class="hl kwa">this</span><span class="hl opt">.</span>scalate<span class="hl opt">) {</span>
        output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">(</span><span class="hl str">':javascript'</span><span class="hl opt">);</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>writer<span class="hl opt">.</span><span class="hl kwd">writeTextContent</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">);</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">(</span><span class="hl str">&quot;&quot;</span> <span class="hl opt">+</span> tagHead <span class="hl opt">+</span> tagAttr<span class="hl opt">);</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>writer<span class="hl opt">.</span><span class="hl kwd">writeTextContent</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">true</span><span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">true</span><span class="hl opt">);</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    Converter<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>style <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">,</span> tagHead<span class="hl opt">,</span> tagAttr<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span><span class="hl kwa">this</span><span class="hl opt">.</span>scalate<span class="hl opt">) {</span>
        output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">(</span><span class="hl str">':css'</span><span class="hl opt">);</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>writer<span class="hl opt">.</span><span class="hl kwd">writeTextContent</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">);</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        output<span class="hl opt">.</span><span class="hl kwd">writeln</span><span class="hl opt">(</span><span class="hl str">&quot;&quot;</span> <span class="hl opt">+</span> tagHead <span class="hl opt">+</span> tagAttr<span class="hl opt">);</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>writer<span class="hl opt">.</span><span class="hl kwd">writeTextContent</span><span class="hl opt">(</span>node<span class="hl opt">,</span> output<span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">,</span> <span class="hl kwa">true</span><span class="hl opt">,</span> <span class="hl kwa">false</span><span class="hl opt">);</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    <span class="hl kwa">return</span> Converter<span class="hl opt">;</span>

  <span class="hl opt">})();</span>

  Output <span class="hl opt">= (</span><span class="hl kwa">function</span><span class="hl opt">() {</span>

    <span class="hl kwa">function</span> <span class="hl kwd">Output</span><span class="hl opt">() {</span>
      <span class="hl kwa">this</span><span class="hl opt">.</span>indents <span class="hl opt">=</span> <span class="hl str">''</span><span class="hl opt">;</span>
    <span class="hl opt">}</span>

    Output<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>enter <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">() {</span>
      <span class="hl kwa">return this</span><span class="hl opt">.</span>indents <span class="hl opt">+=</span> <span class="hl str">'  '</span><span class="hl opt">;</span>
    <span class="hl opt">};</span>

    Output<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>leave <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">() {</span>
      <span class="hl kwa">return this</span><span class="hl opt">.</span>indents <span class="hl opt">=</span> <span class="hl kwa">this</span><span class="hl opt">.</span>indents<span class="hl opt">.</span><span class="hl kwd">substring</span><span class="hl opt">(</span><span class="hl num">2</span><span class="hl opt">);</span>
    <span class="hl opt">};</span>

    Output<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>write <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>data<span class="hl opt">,</span> indent<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        indent <span class="hl opt">=</span> <span class="hl kwa">true</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    Output<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>writeln <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>data<span class="hl opt">,</span> indent<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        indent <span class="hl opt">=</span> <span class="hl kwa">true</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    <span class="hl kwa">return</span> Output<span class="hl opt">;</span>

  <span class="hl opt">})();</span>

  StringOutput <span class="hl opt">= (</span><span class="hl kwa">function</span><span class="hl opt">(</span>_super<span class="hl opt">) {</span>

    <span class="hl kwd">__extends</span><span class="hl opt">(</span>StringOutput<span class="hl opt">,</span> _super<span class="hl opt">);</span>

    <span class="hl kwa">function</span> <span class="hl kwd">StringOutput</span><span class="hl opt">() {</span>
      StringOutput<span class="hl opt">.</span>__super__<span class="hl opt">.</span>constructor<span class="hl opt">.</span><span class="hl kwd">apply</span><span class="hl opt">(</span><span class="hl kwa">this</span><span class="hl opt">,</span> arguments<span class="hl opt">);</span>
      <span class="hl kwa">this</span><span class="hl opt">.</span>fragments <span class="hl opt">= [];</span>
    <span class="hl opt">}</span>

    StringOutput<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>write <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>data<span class="hl opt">,</span> indent<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        indent <span class="hl opt">=</span> <span class="hl kwa">true</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>data <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        data <span class="hl opt">=</span> <span class="hl str">''</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent<span class="hl opt">) {</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>fragments<span class="hl opt">.</span><span class="hl kwd">push</span><span class="hl opt">(</span><span class="hl kwa">this</span><span class="hl opt">.</span>indents <span class="hl opt">+</span> data<span class="hl opt">);</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>fragments<span class="hl opt">.</span><span class="hl kwd">push</span><span class="hl opt">(</span>data<span class="hl opt">);</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    StringOutput<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>writeln <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>data<span class="hl opt">,</span> indent<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        indent <span class="hl opt">=</span> <span class="hl kwa">true</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>data <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        data <span class="hl opt">=</span> <span class="hl str">''</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent<span class="hl opt">) {</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>fragments<span class="hl opt">.</span><span class="hl kwd">push</span><span class="hl opt">(</span><span class="hl kwa">this</span><span class="hl opt">.</span>indents <span class="hl opt">+</span> data <span class="hl opt">+</span> <span class="hl str">'</span><span class="hl esc">\n</span><span class="hl str">'</span><span class="hl opt">);</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>fragments<span class="hl opt">.</span><span class="hl kwd">push</span><span class="hl opt">(</span>data <span class="hl opt">+</span> <span class="hl str">'</span><span class="hl esc">\n</span><span class="hl str">'</span><span class="hl opt">);</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    StringOutput<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>final <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">() {</span>
      <span class="hl kwa">var</span> result<span class="hl opt">;</span>
      result <span class="hl opt">=</span> <span class="hl kwa">this</span><span class="hl opt">.</span>fragments<span class="hl opt">.</span><span class="hl kwd">join</span><span class="hl opt">(</span><span class="hl str">''</span><span class="hl opt">);</span>
      <span class="hl kwa">this</span><span class="hl opt">.</span>fragments <span class="hl opt">= [];</span>
      <span class="hl kwa">return</span> result<span class="hl opt">;</span>
    <span class="hl opt">};</span>

    <span class="hl kwa">return</span> StringOutput<span class="hl opt">;</span>

  <span class="hl opt">})(</span>Output<span class="hl opt">);</span>

  StreamOutput <span class="hl opt">= (</span><span class="hl kwa">function</span><span class="hl opt">(</span>_super<span class="hl opt">) {</span>

    <span class="hl kwd">__extends</span><span class="hl opt">(</span>StreamOutput<span class="hl opt">,</span> _super<span class="hl opt">);</span>

    <span class="hl kwa">function</span> <span class="hl kwd">StreamOutput</span><span class="hl opt">(</span>stream<span class="hl opt">) {</span>
      <span class="hl kwa">this</span><span class="hl opt">.</span>stream <span class="hl opt">=</span> stream<span class="hl opt">;</span>
      StreamOutput<span class="hl opt">.</span>__super__<span class="hl opt">.</span>constructor<span class="hl opt">.</span><span class="hl kwd">apply</span><span class="hl opt">(</span><span class="hl kwa">this</span><span class="hl opt">,</span> arguments<span class="hl opt">);</span>
    <span class="hl opt">}</span>

    StreamOutput<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>write <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>data<span class="hl opt">,</span> indent<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        indent <span class="hl opt">=</span> <span class="hl kwa">true</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>data <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        data <span class="hl opt">=</span> <span class="hl str">''</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent<span class="hl opt">) {</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>stream<span class="hl opt">.</span><span class="hl kwd">write</span><span class="hl opt">(</span><span class="hl kwa">this</span><span class="hl opt">.</span>indents <span class="hl opt">+</span> data<span class="hl opt">);</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>stream<span class="hl opt">.</span><span class="hl kwd">write</span><span class="hl opt">(</span>data<span class="hl opt">);</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    StreamOutput<span class="hl opt">.</span><span class="hl kwa">prototype</span><span class="hl opt">.</span>writeln <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>data<span class="hl opt">,</span> indent<span class="hl opt">) {</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        indent <span class="hl opt">=</span> <span class="hl kwa">true</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>data <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        data <span class="hl opt">=</span> <span class="hl str">''</span><span class="hl opt">;</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>indent<span class="hl opt">) {</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>stream<span class="hl opt">.</span><span class="hl kwd">write</span><span class="hl opt">(</span><span class="hl kwa">this</span><span class="hl opt">.</span>indents <span class="hl opt">+</span> data <span class="hl opt">+</span> <span class="hl str">'</span><span class="hl esc">\n</span><span class="hl str">'</span><span class="hl opt">);</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        <span class="hl kwa">return this</span><span class="hl opt">.</span>stream<span class="hl opt">.</span><span class="hl kwd">write</span><span class="hl opt">(</span>data <span class="hl opt">+</span> <span class="hl str">'</span><span class="hl esc">\n</span><span class="hl str">'</span><span class="hl opt">);</span>
      <span class="hl opt">}</span>
    <span class="hl opt">};</span>

    <span class="hl kwa">return</span> StreamOutput<span class="hl opt">;</span>

  <span class="hl opt">})(</span>Output<span class="hl opt">);</span>

  scope<span class="hl opt">.</span>Output <span class="hl opt">=</span> Output<span class="hl opt">;</span>

  scope<span class="hl opt">.</span>StringOutput <span class="hl opt">=</span> StringOutput<span class="hl opt">;</span>

  scope<span class="hl opt">.</span>Converter <span class="hl opt">=</span> Converter<span class="hl opt">;</span>

  scope<span class="hl opt">.</span>Writer <span class="hl opt">=</span> Writer<span class="hl opt">;</span>

  <span class="hl kwa">if</span> <span class="hl opt">(</span><span class="hl kwa">typeof</span> exports <span class="hl opt">!==</span> <span class="hl str">&quot;undefined&quot;</span> <span class="hl opt">&amp;&amp;</span> exports <span class="hl opt">!==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
    scope<span class="hl opt">.</span>Parser <span class="hl opt">=</span> Parser<span class="hl opt">;</span>
    scope<span class="hl opt">.</span>StreamOutput <span class="hl opt">=</span> StreamOutput<span class="hl opt">;</span>
    scope<span class="hl opt">.</span>convert <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>input<span class="hl opt">,</span> output<span class="hl opt">,</span> options<span class="hl opt">) {</span>
      <span class="hl kwa">var</span> _ref1<span class="hl opt">;</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>options <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        options <span class="hl opt">= {};</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">if</span> <span class="hl opt">((</span>_ref1 <span class="hl opt">=</span> options<span class="hl opt">.</span>parser<span class="hl opt">) ==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
        options<span class="hl opt">.</span>parser <span class="hl opt">=</span> <span class="hl kwa">new</span> <span class="hl kwd">Parser</span><span class="hl opt">(</span>options<span class="hl opt">);</span>
      <span class="hl opt">}</span>
      <span class="hl kwa">return</span> options<span class="hl opt">.</span>parser<span class="hl opt">.</span><span class="hl kwd">parse</span><span class="hl opt">(</span>input<span class="hl opt">,</span> <span class="hl kwa">function</span><span class="hl opt">(</span>errors<span class="hl opt">,</span> window<span class="hl opt">) {</span>
        <span class="hl kwa">var</span> _ref2<span class="hl opt">;</span>
        <span class="hl kwa">if</span> <span class="hl opt">(</span>errors <span class="hl opt">!=</span> <span class="hl kwa">null</span> ? errors<span class="hl opt">.</span>length <span class="hl opt">:</span> <span class="hl kwa">void</span> <span class="hl num">0</span><span class="hl opt">) {</span>
          <span class="hl kwa">return</span> errors<span class="hl opt">;</span>
        <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
          <span class="hl kwa">if</span> <span class="hl opt">(</span>output <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
            output <span class="hl opt">=</span> <span class="hl kwa">new</span> <span class="hl kwd">StreamOutput</span><span class="hl opt">(</span>process<span class="hl opt">.</span>stdout<span class="hl opt">);</span>
          <span class="hl opt">}</span>
          <span class="hl kwa">if</span> <span class="hl opt">((</span>_ref2 <span class="hl opt">=</span> options<span class="hl opt">.</span>converter<span class="hl opt">) ==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
            options<span class="hl opt">.</span>converter <span class="hl opt">=</span> <span class="hl kwa">new</span> <span class="hl kwd">Converter</span><span class="hl opt">(</span>options<span class="hl opt">);</span>
          <span class="hl opt">}</span>
          <span class="hl kwa">return</span> options<span class="hl opt">.</span>converter<span class="hl opt">.</span><span class="hl kwd">document</span><span class="hl opt">(</span>window<span class="hl opt">.</span>document<span class="hl opt">,</span> output<span class="hl opt">);</span>
        <span class="hl opt">}</span>
      <span class="hl opt">});</span>
    <span class="hl opt">};</span>
  <span class="hl opt">}</span>

  scope<span class="hl opt">.</span>convertHtml <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>html<span class="hl opt">,</span> options<span class="hl opt">,</span> cb<span class="hl opt">) {</span>
    <span class="hl kwa">var</span> _ref1<span class="hl opt">;</span>
    <span class="hl kwa">if</span> <span class="hl opt">(</span>options <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
      options <span class="hl opt">= {};</span>
    <span class="hl opt">}</span>
    <span class="hl kwa">if</span> <span class="hl opt">((</span>_ref1 <span class="hl opt">=</span> options<span class="hl opt">.</span>parser<span class="hl opt">) ==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
      options<span class="hl opt">.</span>parser <span class="hl opt">=</span> <span class="hl kwa">new</span> <span class="hl kwd">Parser</span><span class="hl opt">(</span>options<span class="hl opt">);</span>
    <span class="hl opt">}</span>
    <span class="hl kwa">return</span> options<span class="hl opt">.</span>parser<span class="hl opt">.</span><span class="hl kwd">parse</span><span class="hl opt">(</span>html<span class="hl opt">,</span> <span class="hl kwa">function</span><span class="hl opt">(</span>errors<span class="hl opt">,</span> window<span class="hl opt">) {</span>
      <span class="hl kwa">var</span> output<span class="hl opt">,</span> _ref2<span class="hl opt">,</span> _ref3<span class="hl opt">;</span>
      <span class="hl kwa">if</span> <span class="hl opt">(</span>errors <span class="hl opt">!=</span> <span class="hl kwa">null</span> ? errors<span class="hl opt">.</span>length <span class="hl opt">:</span> <span class="hl kwa">void</span> <span class="hl num">0</span><span class="hl opt">) {</span>
        <span class="hl kwa">return</span> errors<span class="hl opt">;</span>
      <span class="hl opt">}</span> <span class="hl kwa">else</span> <span class="hl opt">{</span>
        output <span class="hl opt">= (</span>_ref2 <span class="hl opt">=</span> options<span class="hl opt">.</span>output<span class="hl opt">) !=</span> <span class="hl kwa">null</span> ? _ref2 <span class="hl opt">:</span> <span class="hl kwa">new</span> <span class="hl kwd">StringOutput</span><span class="hl opt">();</span>
        <span class="hl kwa">if</span> <span class="hl opt">((</span>_ref3 <span class="hl opt">=</span> options<span class="hl opt">.</span>converter<span class="hl opt">) ==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
          options<span class="hl opt">.</span>converter <span class="hl opt">=</span> <span class="hl kwa">new</span> <span class="hl kwd">Converter</span><span class="hl opt">(</span>options<span class="hl opt">);</span>
        <span class="hl opt">}</span>
        options<span class="hl opt">.</span>converter<span class="hl opt">.</span><span class="hl kwd">document</span><span class="hl opt">(</span>window<span class="hl opt">.</span>document<span class="hl opt">,</span> output<span class="hl opt">);</span>
        <span class="hl kwa">if</span> <span class="hl opt">(</span>cb <span class="hl opt">!=</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
          <span class="hl kwa">return</span> <span class="hl kwd">cb</span><span class="hl opt">(</span><span class="hl kwa">null</span><span class="hl opt">,</span> output<span class="hl opt">.</span><span class="hl kwd">final</span><span class="hl opt">());</span>
        <span class="hl opt">}</span>
      <span class="hl opt">}</span>
    <span class="hl opt">});</span>
  <span class="hl opt">};</span>

  scope<span class="hl opt">.</span>convertDocument <span class="hl opt">=</span> <span class="hl kwa">function</span><span class="hl opt">(</span>document<span class="hl opt">,</span> options<span class="hl opt">,</span> cb<span class="hl opt">) {</span>
    <span class="hl kwa">var</span> output<span class="hl opt">,</span> _ref1<span class="hl opt">,</span> _ref2<span class="hl opt">;</span>
    <span class="hl kwa">if</span> <span class="hl opt">(</span>options <span class="hl opt">==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
      options <span class="hl opt">= {};</span>
    <span class="hl opt">}</span>
    output <span class="hl opt">= (</span>_ref1 <span class="hl opt">=</span> options<span class="hl opt">.</span>output<span class="hl opt">) !=</span> <span class="hl kwa">null</span> ? _ref1 <span class="hl opt">:</span> <span class="hl kwa">new</span> <span class="hl kwd">StringOutput</span><span class="hl opt">();</span>
    <span class="hl kwa">if</span> <span class="hl opt">((</span>_ref2 <span class="hl opt">=</span> options<span class="hl opt">.</span>converter<span class="hl opt">) ==</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
      options<span class="hl opt">.</span>converter <span class="hl opt">=</span> <span class="hl kwa">new</span> <span class="hl kwd">Converter</span><span class="hl opt">(</span>options<span class="hl opt">);</span>
    <span class="hl opt">}</span>
    options<span class="hl opt">.</span>converter<span class="hl opt">.</span><span class="hl kwd">document</span><span class="hl opt">(</span>document<span class="hl opt">,</span> output<span class="hl opt">);</span>
    <span class="hl kwa">if</span> <span class="hl opt">(</span>cb <span class="hl opt">!=</span> <span class="hl kwa">null</span><span class="hl opt">) {</span>
      <span class="hl kwa">return</span> <span class="hl kwd">cb</span><span class="hl opt">(</span><span class="hl kwa">null</span><span class="hl opt">,</span> output<span class="hl opt">.</span><span class="hl kwd">final</span><span class="hl opt">());</span>
    <span class="hl opt">}</span>
  <span class="hl opt">};</span>

<span class="hl opt">}).</span><span class="hl kwd">call</span><span class="hl opt">(</span><span class="hl kwa">this</span><span class="hl opt">);</span>
</pre>
</body>
</html>
<!--HTML generated by highlight 3.9, http://www.andre-simon.de/-->`,
			ExpectedJade: `doctype html
html
  head
    meta(http-equiv='content-type', content='text/html; charset=ISO-8859-1')
    title html2jade.js
    link(rel='stylesheet', type='text/css', href='highlight.css')
  body.hl
    pre.hl.
      \n      
               Converter       Output       Parser       StreamOutput       StringOutput       Writer       publicIdDocTypeNames       scope       systemIdDocTypeNames       _ref      
          __hasProp       hasOwnProperty      
          __extends              child       parent                     key        parent                    __hasProp      parent       key       child              parent                           constructor        child       ctor              parent       child                            child      __super__        parent              child      
      
        scope               exports                      exports               ? exports       _ref              Html2Jade              ? _ref              Html2Jade       
      
        Parser       
      
                       options      
                         options              
              options       
                  
                  jsdom              
                
      
          Parser      parse              arg       cb      
                         arg      
                           
                                
                    jsdom      arg       cb      
                  
                
      
                 Parser      
      
              
      
        Writer       
      
                       options      
                   _ref1       _ref2      
                         options              
              options       
                  
                  wrapLength       _ref1        options      wrapLength              ? _ref1              
                  scalate       _ref2        options      scalate              ? _ref2              
                  attrSep              scalate ?                     
                
      
          Writer      tagHead              node      
                   classes       result      
            result        node      tagName               ? node      tagName             
                         node      id      
              result                      node      id      
                  
                         node       node      length              
              classes        node      \s      item      
                             item               item      length              
                    
              result                      classes      
                  
                         result      length              
              result              
                  
                   result      
                
      
          Writer      tagAttr              node      
                   attr       attrs       nodeName       result       _i       _len      
            attrs        node      attributes      
                         attrs || attrs      length              
                           
                                
              result       
                           _i               _len        attrs      length       _i        _len       _i      
                attr        attrs      
                             attr       nodeName        attr      nodeName      
                               nodeName                      nodeName                                   attr      nodeValue              
                    result      attr      nodeName                      attr      nodeValue      
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            |       
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
                   
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            html      
            
            
            
                   
            
            
            
            
            
            html      
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            script      style      
            src      
            
            
            script      
            
            style      
            
            
            conditional      
            
                    node       output      
                                tagName      
              output      tagHead        tagAttr              
              output      
              firstline              
                    writer      node             child      
                       data      
                             child      nodeType              
                  data        child      data      
                               data               data      length              
                                 firstline      
                                   data      |      |             
                        data        data      |      |             
                            
                      data                      data      
                      firstline              
                          
                    data        data      g             
                    data        data      |      |      g                     output      indents      
                           output      data      
                        
                      
                    
              output      
                     output      
                                tagText      
                           tagText      length                      tagText             
                tagText                      tagText      
                    
                     output      tagHead        tagAttr                      tagText      
                                
              output      tagHead        tagAttr      
                    node       output      
                  
                
      
          Converter      children              parent       output       indent      
                   _this              
                         indent              
              indent              
                  
                         indent      
              output      
                  
                  writer      parent             child      
                     nodeType      
              nodeType        child      nodeType      
                           nodeType              
                       _this      child       output      
                                  nodeType              
                             parent      _nodeName              
                         _this      child       output                           
                                    
                         _this      child       output      
                      
                                  nodeType              
                       _this      child       output      
                    
                  
                         indent      
                     output      
                  
                
      
          Converter      text              node       output       pipe       trim       wrap      
            node      
                  writer      node       output       pipe       trim       wrap      
                
      
          Converter      comment              node       output      
                   condition       data       lines      
              _this              
            condition        node      data      \s      \      \s      ^\      \      
                         condition      
              data        node      data ||       
                           data      length               || data      |      
                       output             data      
                                  
                output      
                output      
                lines        data      |      
                lines      line      
                         _this      writer      line       output                           
                      
                       output      
                    
                                
                    node       condition       output      
                  
                
      
          Converter      conditional              node       condition       output      
                   conditionalElem       innerHTML      
            innerHTML        node      textContent      \s      \      \s      ^\      \      \s*/                    
                         innerHTML             
              condition                      condition              
              innerHTML              
                  
            conditionalElem        node      ownerDocument      
            conditionalElem       condition      
                         innerHTML      
              conditionalElem      innerHTML        innerHTML      
                  
                   node      parentNode      conditionalElem       node      nextSibling      
                
      
          Converter      script              node       output       tagHead       tagAttr      
                         scalate      
              output      
                    writer      node       output                           
                                
              output              tagHead        tagAttr      
                    writer      node       output                                  
                  
                
      
          Converter      style              node       output       tagHead       tagAttr      
                         scalate      
              output      
                    writer      node       output                           
                                
              output              tagHead        tagAttr      
                    writer      node       output                           
                  
                
      
                 Converter      
      
              
      
        Output       
      
                       
                  indents              
                
      
          Output      enter              
                  indents              
                
      
          Output      leave              
                  indents              indents      
                
      
          Output      write              data       indent      
                         indent              
              indent              
                  
                
      
          Output      writeln              data       indent      
                         indent              
              indent              
                  
                
      
                 Output      
      
              
      
        StringOutput       _super      
      
                StringOutput       _super      
      
                       
            StringOutput      __super__      constructor       arguments      
                  fragments       
                
      
          StringOutput      write              data       indent      
                         indent              
              indent              
                  
                         data              
              data              
                  
                         indent      
                    fragments      indents        data      
                                
                    fragments      data      
                  
                
      
          StringOutput      writeln              data       indent      
                         indent              
              indent              
                  
                         data              
              data              
                  
                         indent      
                    fragments      indents        data              
                                
                    fragments      data              
                  
                
      
          StringOutput      final              
                   result      
            result              fragments      
                  fragments       
                   result      
                
      
                 StringOutput      
      
              Output      
      
        StreamOutput       _super      
      
                StreamOutput       _super      
      
                       stream      
                  stream        stream      
            StreamOutput      __super__      constructor       arguments      
                
      
          StreamOutput      write              data       indent      
                         indent              
              indent              
                  
                         data              
              data              
                  
                         indent      
                    stream      indents        data      
                                
                    stream      data      
                  
                
      
          StreamOutput      writeln              data       indent      
                         indent              
              indent              
                  
                         data              
              data              
                  
                         indent      
                    stream      indents        data              
                                
                    stream      data              
                  
                
      
                 StreamOutput      
      
              Output      
      
        scope      Output        Output      
      
        scope      StringOutput        StringOutput      
      
        scope      Converter        Converter      
      
        scope      Writer        Writer      
      
                      exports                      exports              
          scope      Parser        Parser      
          scope      StreamOutput        StreamOutput      
          scope      convert              input       output       options      
                   _ref1      
                         options              
              options       
                  
                         _ref1        options      parser             
              options      parser                     options      
                  
                   options      parser      input             errors       window      
                     _ref2      
                           errors               ? errors      length                     
                       errors      
                                  
                             output              
                  output                     process      stdout      
                      
                             _ref2        options      converter             
                  options      converter                     options      
                      
                       options      converter      window      document       output      
                    
                  
                
              
      
        scope      convertHtml              html       options       cb      
                 _ref1      
                       options              
            options       
                
                       _ref1        options      parser             
            options      parser                     options      
                
                 options      parser      html             errors       window      
                   output       _ref2       _ref3      
                         errors               ? errors      length                     
                     errors      
                                
              output       _ref2        options      output              ? _ref2                     
                           _ref3        options      converter             
                options      converter                     options      
                    
              options      converter      window      document       output      
                           cb              
                              output      
                    
                  
                
              
      
        scope      convertDocument              document       options       cb      
                 output       _ref1       _ref2      
                       options              
            options       
                
          output       _ref1        options      output              ? _ref1                     
                       _ref2        options      converter             
            options      converter                     options      
                
          options      converter      document       output      
                       cb              
                          output      
                
              
      
            
            
// HTML generated by highlight 3.9, http://www.andre-simon.de/
`,
			NilAssertion: assert.Nil,
		},

		{
			Desc:    "TEST018 - Pre 2",
			Options: defaultOptionsWithHead,
			Skip:    doSKip,
			SourceHTML: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<html>
<head>
<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1">
<title>html2jade.js</title>
<link rel="stylesheet" type="text/css" href="highlight.css">
</head>
<body class="hl">
<pre class="hl"><span class="hl slc">// Generated by CoffeeScript 1.3.3</span>
<span class="hl opt">(</span><span class="hl kwa">function</span><span class="hl opt">() {</span>
  <span class="hl kwa">var</span> Converter<span class="hl opt">,</span> Output<span class="hl opt">,</span> Parser<span class="hl opt">,</span> StreamOutput<span class="hl opt">,</span> StringOutput<span class="hl opt">,</span> Writer<span class="hl opt">,</span> publicIdDocTypeNames<span class="hl opt">,</span> scope<span class="hl opt">,</span> systemIdDocTypeNames<span class="hl opt">,</span> _ref<span class="hl opt">,</span>
</pre>
`,
			ExpectedJade: `doctype html
html
  head
    meta(http-equiv='content-type', content='text/html; charset=ISO-8859-1')
    title html2jade.js
    link(rel='stylesheet', type='text/css', href='highlight.css')
  body.hl
    pre.hl.
      \n      
               Converter       Output       Parser       StreamOutput       StringOutput       Writer       publicIdDocTypeNames       scope       systemIdDocTypeNames       _ref      
            
`,
			NilAssertion: assert.Nil,
		},

		{
			Desc:    "TEST019 - Pre 3",
			Options: defaultOptionsWithHead,
			Skip:    doSKip,
			SourceHTML: `<!DOCTYPE html><html><head><meta http-equiv="content-type" content="text/html; charset=ISO-8859-1"><title>html2jade.js</title><link rel="stylesheet" type="text/css" href="highlight.css"></head><body class="hl"><pre class="hl"><span class="hl slc">// Generated by CoffeeScript 1.3.3
</span><span class="hl opt">(</span><span class="hl kwa">function</span><span class="hl opt">() {</span><span class="hl kwa">var Converter</span><span class="hl opt">,
Output</span><span class="hl opt">,</span>Parser<span class="hl opt">,</span>StreamOutput<span class="hl opt">,</span>StringOutput<span class="hl opt">,</span>Writer<span class="hl opt">,</span>publicIdDocTypeNames<span class="hl opt">,</span>scope<span class="hl opt">,</span>systemIdDocTypeNames<span class="hl opt">,</span>_ref<span class="hl opt">,</span></pre></body></html>`,
			ExpectedJade: `doctype html
html
  head
    meta(http-equiv='content-type', content='text/html; charset=ISO-8859-1')
    title html2jade.js
    link(rel='stylesheet', type='text/css', href='highlight.css')
  body.hl
    pre.hl.
      \nParser      StreamOutput      StringOutput      Writer      publicIdDocTypeNames      scope      systemIdDocTypeNames      _ref      
`,
			NilAssertion: assert.Nil,
		},

		{
			Desc:    "TEST020 - Test",
			Skip:    doSKip,
			Options: defaultOptions,
			SourceHTML: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<html>
<head>
<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1">
<title>html2jade.js</title>
<link rel="stylesheet" type="text/css" href="highlight.css">
</head>
<body class="hl">
<pre class="hl"><span class="hl slc">// Generated by CoffeeScript 1.3.3</span>
<span class="hl opt">(</span><span class="hl kwa">function</span><span class="hl opt">() {</span>
  <span class="hl kwa">var</span> Converter<span class="hl opt">,</span> Output<span class="hl opt">,</span> Parser<span class="hl opt">,</span> StreamOutput<span class="hl opt">,</span> StringOutput<span class="hl opt">,</span> Writer<span class="hl opt">,</span> publicIdDocTypeNames<span class="hl opt">,</span> scope<span class="hl opt">,</span> systemIdDocTypeNames<span class="hl opt">,</span> _ref<span class="hl opt">,</span>
</pre>
`,
			ExpectedJade: `doctype html
html
  head
    meta(http-equiv='content-type', content='text/html; charset=ISO-8859-1')
    title html2jade.js
    link(rel='stylesheet', type='text/css', href='highlight.css')
  body.hl
    pre.hl.
      \n      
               Converter       Output       Parser       StreamOutput       StringOutput       Writer       publicIdDocTypeNames       scope       systemIdDocTypeNames       _ref      
            
`,
			NilAssertion: assert.Nil,
		},

		{
			Desc:    "TEST021 - TextArea Javascript",
			Options: defaultOptions,
			Skip:    doSKip,
			SourceHTML: `<textarea id="text-area">javascript:window.s=document.createElement('script');window.sc=document.getElementsByTagName("body")[0]||document.getElementsByTagName("head")[0];s.src="http://xyz.com/path/app.js";sc.appendChild(s)
</textarea>
`,
			ExpectedJade: `html
  body
    textarea#text-area
      | javascript:window.s=document.createElement(&apos;script&apos;);window.sc=document.getElementsByTagName(&quot;body&quot;)[0]||document.getElementsByTagName(&quot;head&quot;)[0];s.src=&quot;http://xyz.com/path/app.js&quot;;sc.appendChild(s)
`,
			NilAssertion: assert.Nil,
		},

		{
			Desc:    "TEST022 - Whitespace",
			Options: defaultOptions,
			Skip:    doSKip,
			SourceHTML: `<p>Here is a <a href="#">link</a> with whitespaces around it</p>
`,
			ExpectedJade: `html
  body
    p
      | Here is a 
      a(href='#') link
      |  with whitespaces around it
`,
			NilAssertion: assert.Nil,
		},

		{
			Desc:    "TEST023 - Whitespace 2",
			Options: defaultOptions,
			Skip:    doSKip,
			SourceHTML: `<p>Hey there, <a href="#">html2jade</a> <strong>is awesome</strong></p>
`,
			ExpectedJade: `html
  body
    p
      | Hey there, 
      a(href='#') html2jade
      |  
      strong is awesome
`,
			NilAssertion: assert.Nil,
		},
	}

	for _, tc := range testCases {
		logrus.Infoln("test case " + tc.Desc)

		tc := tc
		if tc.Skip {
			logrus.Warnln("skipping " + tc.Desc)
			// t.Skip(tc.Desc)
			continue
		}

		t.Run(tc.Desc, func(t *testing.T) {

			jadeConvertor := pkg.NewHtml2PugConvertor(tc.Options)
			// jadeConvertor.Convert(tc.SourceHTML, (*entities.IStringWriter)(&stringOutput), tc.Options)

			callback := func(err error, jadeOutput string) {
				assert.Equal(t, tc.ExpectedJade, jadeOutput)
			}
			jadeConvertor.ConvertHTML(tc.SourceHTML, (entities.Html2JadeConvertorConvertDocumentCallback)(callback))

		})
	}
}
