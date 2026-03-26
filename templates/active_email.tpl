<!-- index.tmpl -->
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
</head>
<body style="margin:0;background-color: #005db4;">
    <div style="width: 100%;height:100vh ;display: flex;flex:1;align-items: center;justify-content: center;flex-direction: column;">
        <div>
            <p style="font-size:24px; color:#fff; padding: 3px 16px;">
            <span>{{.content}}</span>
            </p>
        </div>
        <div style="width: 100%;align-items: center;display: flex;justify-content: center;flex-direction: row;">
            <p style="font-size:16px; color:#999; padding: 3px 16px;">辽宁业聘科技有限公司</p>
        </div>
        
    </div>
    
</body>
</html>