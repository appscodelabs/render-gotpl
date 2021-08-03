{{ range .providers }}

        <!-- single features start  -->
        <div class="single-features">
          <div class="features-name"><a href="https://github.com/kubeform/provider-{{ . }}-api">Terraform {{ . | title }} Provider</a></div>
          <!-- features plan wrapper  -->
          <div class="features-plan-wrapper">
            <!-- community plan start  -->
            <div class="features-plan community"><span class="t-icon"><img src="/assets/images/icon/check-icon.svg"
                  alt="check-icon"></span> </div>
            <!-- community plan end  -->
            <!-- enterprise plan start  -->
            <div class="features-plan enterprise">
              <div class="e-plan"><span class="t-icon"><img src="/assets/images/icon/check-icon.svg"
                    alt="check-icon"></span></div>
              <div class="e-plan"><span class="t-icon"><img src="/assets/images/icon/check-icon.svg"
                    alt="check-icon"></span></div>
              <div class="e-plan"><span class="t-icon"><img src="/assets/images/icon/check-icon.svg"
                    alt="check-icon"></span></div>
            </div>
            <!-- enterprise plan end  -->
          </div>
          <!-- features plan wrapper  -->
        </div>
        <!-- single features end -->
{{ end }}