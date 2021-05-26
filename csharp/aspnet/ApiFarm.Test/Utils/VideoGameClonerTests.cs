using System;
using System.Collections.Generic;
using ApiFarm.Models.Impl;
using ApiFarm.Utils.Impl;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Utils
{
    public class VideoGameClonerTests
    {
        private VideoGameCloner subject;

        [SetUp]
        public void SetUp()
        {
            subject = new VideoGameCloner();
        }

        private class CloneShould : VideoGameClonerTests
        {
            [Test]
            public void ReturnDeepCloneWithDifferentReferences()
            {
                // Arrange
                var original = new VideoGame(5)
                {
                    Name = "Vikings at Sea II",
                    Developers = new List<string> { "A" },
                    Publishers = new List<string> { "A", "B" },
                    Directors = new List<string> { "A" },
                    Producers = new List<string> { "A" },
                    Designers = new List<string> { "A", "B", "C" },
                    Programmers = new List<string> { "A", "B", "C", "D" },
                    Artists = new List<string> { "A", "B", "C" },
                    Composers = new List<string> { "A", "B" },
                    Platforms = new List<string> { "A", "B", "C" },
                    DateReleased = DateTime.Now,
                };

                // Act
                var clone = subject.Clone(original);

                // Assert
                clone.ShouldBeEquivalentTo(original);

                clone.ShouldNotBe(original);
                clone.Developers.ShouldNotBe(original.Developers);
                clone.Publishers.ShouldNotBe(original.Publishers);
                clone.Directors.ShouldNotBe(original.Directors);
                clone.Producers.ShouldNotBe(original.Producers);
                clone.Designers.ShouldNotBe(original.Designers);
                clone.Programmers.ShouldNotBe(original.Programmers);
                clone.Artists.ShouldNotBe(original.Artists);
                clone.Composers.ShouldNotBe(original.Composers);
                clone.Platforms.ShouldNotBe(original.Platforms);
            }
        }
    }
}
