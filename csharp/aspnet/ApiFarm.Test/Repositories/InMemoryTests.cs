using System.Collections.Generic;
using ApiFarm.Models;
using ApiFarm.Repositories;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Repositories
{
    public class InMemoryTests
    {
        private InMemory<Model> subject;

        private List<Model> models;

        [SetUp]
        public void Setup()
        {
            models = new List<Model>();
            subject = new InMemory<Model>(models);
        }

        private class GetAllShould : InMemoryTests
        {
            [Test]
            public void ReturnEmptyModelsList()
            {
                // Arrange
                subject = new InMemory<Model>();

                // Act
                var retrievedModels = subject.GetAll();

                // Assert
                retrievedModels.ShouldBeOfType<List<Model>>();
                retrievedModels.ShouldBeEmpty();
            }

            [Test]
            public void ReturnStubbedModelsList()
            {
                // Act
                var retrievedModels = subject.GetAll();

                // Assert
                retrievedModels.ShouldBe(models);
            }
        }

        private class Model : IModel
        {
            public uint Id { get; set; }
        }
    }
}
