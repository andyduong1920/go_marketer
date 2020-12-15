<h1>Register</h1>

<form action="/register" method="post">
  <input name="_csrf_token" type="hidden" value="AzY2GiIhWggpEFRiImY9L1EXM3oJI2FzNrOlMP00dAeOeVtI4NuLPgY1" />

  <div class="form-group">
    <label for="user_email">Email</label>
    <input class="form-control" id="user_email" name="email" type="email" required="" />
  </div>

  <div class="form-group">
    <label for="user_password">Password</label>
    <input class="form-control" id="user_password" name="password" type="password" required="" />
  </div>

  <div>
    <button class="btn btn-primary" type="submit">Register</button>
  </div>
</form>

<div class="mt-3">
  <a href="/login">Login</a> |
  <a href="/reset_password">Forgot your password?</a>
</div>
