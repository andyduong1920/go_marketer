<h1>Register</h1>

<form action="/register" method="post">
  <div class="form-group">
    <label for="user_email">Email</label>
    <input class="form-control" id="user_email" name="email" type="email" required="true" />
  </div>

  <div class="form-group">
    <label for="user_password">Password</label>
    <input class="form-control" id="user_password" name="password" type="password" required="true" />
  </div>

  <div>
    <button class="btn btn-primary mt-3" type="submit">Register</button>
  </div>
</form>
