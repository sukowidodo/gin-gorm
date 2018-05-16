<html>
<head>
    <title>
        Form User
    </title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.0/umd/popper.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/js/bootstrap.min.js"></script>
</head>
<body>
    <div class="container">
            <h1>
                {{ .title }}
            </h1>
            <form action="/api/v1/users" method="post">
                <div class="form-group">
                    <label for="email">FirsName:</label>
                    <input type="text" class="form-control" id="email">
                </div>
                <div class="form-group">
                    <label for="pwd">LastName:</label>
                    <input type="text" class="form-control" id="pwd">
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form> 
    </div>
</body>
</html>