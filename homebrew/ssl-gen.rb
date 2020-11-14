require 'rbconfig'
class Ssl-Gen < Formula
  desc ""
  homepage "https://github.com/ashsmith/ssl-gen"
  version "1.0.0"

  if Hardware::CPU.is_64_bit?
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/ashsmith/ssl-gen/releases/download/v1.0.0/ssl-gen_1.0.0_darwin_amd64.zip"
      sha256 "a9dd1256cc5955be308f0ee310c04173f8b0c259a30c56f08e03835bdc6789fc"
    when /linux/
      url "https://github.com/ashsmith/ssl-gen/releases/download/v1.0.0/ssl-gen_1.0.0_linux_amd64.tar.gz"
      sha256 "c3388961535c5db017e46a059485ac9d4eab52bba02182fd240bf91807891884"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  else
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/ashsmith/ssl-gen/releases/download/v1.0.0/ssl-gen_1.0.0_darwin_386.zip"
      sha256 "72b722170f5524d2337f655ba6b2119028dbc3fcf8b352784511345f35f40ac6"
    when /linux/
      url "https://github.com/ashsmith/ssl-gen/releases/download/v1.0.0/ssl-gen_1.0.0_linux_386.tar.gz"
      sha256 "3f6ee41a15934207259b042d35dbb089ae21c572c5ca6d5aa5829d45e9d8d250"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  end

  def install
    bin.install "ssl-gen"
  end

  test do
    system "ssl-gen"
  end

end
