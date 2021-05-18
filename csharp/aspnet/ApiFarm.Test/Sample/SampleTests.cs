namespace Tests.SampleTests
{
    using ApiFarm.Sample;
    using NUnit.Framework;
    using Shouldly;

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