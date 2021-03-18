require_relative "../lib/sample"

describe "checks the addition of two files" do
    it "adds two numbers together" do
        expect(add(3, 5)).to be(8)
    end
end