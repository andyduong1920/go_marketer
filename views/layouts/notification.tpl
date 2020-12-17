<div class="notification">
  {{if .flash.notice}}
    <p class="alert alert-info my-3" role="alert">
      {{.flash.notice}}
    </p>
  {{end}}
  
  {{if .flash.error}}
    <p class="alert alert-danger my-3" role="alert">
      {{.flash.error}}
    </p>
  {{end}}
</div>
