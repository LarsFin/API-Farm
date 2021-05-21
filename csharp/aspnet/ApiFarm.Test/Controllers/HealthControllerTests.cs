using ApiFarm.Controllers;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Controllers
{
    public class HealthControllerTests
    {
        private HealthController subject;

        [SetUp]
        protected void SetUp()
        {
            subject = new HealthController();
        }

        private class PingShould : HealthControllerTests
        {
            [Test]
            public void ReturnPong()
            {
                // Arrange
                var expected = "pong";

                // Act
                var actual = subject.Ping();

                // Assert
                actual.ShouldBe(expected);
            }
        }
    }
}
