using ApiFarm.test_sample;
using NUnit.Framework;
using Shouldly;

namespace Tests.SampleTests
{
    public class SampleTests
    {
        [Test]
        public void Test1()
        {
            // Arrange
            const int expected = 20;

            // Act
            var actual = Sample.Add(12, 8);

            // Assert
            actual.ShouldBe(expected);
        }
    }
}