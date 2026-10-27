require "json"

class AguastoResearch < Formula
  desc "Aguasto public research and educational calculator CLI"
  homepage "https://www.aguasto.com/developers"
  url "https://github.com/Altssdata/aguasto-agents/releases/download/v1.0.0/aguasto-research-js.tgz"
  version "1.0.0"
  sha256 "1627b8b43c2f7c4dd51454880a410b287447f422bb539805237bc8fbe746b7d3"
  license :cannot_represent

  depends_on "node"

  def install
    source = (buildpath/"package/cli.mjs").exist? ? buildpath/"package" : buildpath
    libexec.install source.children
    chmod 0755, libexec/"cli.mjs"
    (bin/"aguasto-research").write <<~SH
      #!/bin/sh
      exec "#{Formula["node"].opt_bin}/node" "#{libexec}/cli.mjs" "$@"
    SH
  end

  test do
    result = JSON.parse(shell_output("#{bin}/aguasto-research fixtures"))
    assert_equal true, result.fetch("offline_fixture")
    assert_equal 35, result.fetch("weight_memory").fetch("result").fetch("gb")
  end
end
