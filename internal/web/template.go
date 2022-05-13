package web

var indexPageContent = `
<html>
<head>
    <title>Index</title>
    <style>
        .profile-name{
            display:inline-block;
            width:6rem;
        }
    </style>
</head>
<body>
Hydra Home Page<br>
<br>
<table>
    <tr>
        <td><a href=/debug/pprof>Pprof</a></td>
    </tr>
    <tr><td><br></td></tr>

    <tr>
        <td><a href=/debug/metrics>Metrics</a></td>
    </tr>

</table>
<br/>
</body>
</html>
`
